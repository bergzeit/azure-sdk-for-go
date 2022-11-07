//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package shared

import (
	"context"
	"errors"
	"sync"
)

// BatchTransferOptions identifies options used by doBatchTransfer.
type BatchTransferOptions struct {
	TransferSize  int64
	ChunkSize     int64
	Concurrency   uint16
	Operation     func(offset int64, chunkSize int64, ctx context.Context) error
	OperationName string
}

// DoBatchTransfer helps to execute operations in a batch manner.
// Can be used by users to customize batch works (for other scenarios that the SDK does not provide)
func DoBatchTransfer(ctx context.Context, o *BatchTransferOptions) error {
	if o.ChunkSize == 0 {
		return errors.New("ChunkSize cannot be 0")
	}

	if o.Concurrency == 0 {
		o.Concurrency = 5 // default concurrency
	}

	// Prepare and do parallel operations.
	numChunks := uint16(((o.TransferSize - 1) / o.ChunkSize) + 1)
	operationChannel := make(chan func() error, o.Concurrency) // Create the channel that release 'concurrency' goroutines concurrently
	operationResponseChannel := make(chan error)               // Listen for operation responses and react to them
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a background listener to the error channel. As soon as one error
	// is returned from an operation cancel all remaining operations and return
	// this first error.
	var firstErr error = nil
	errWg := sync.WaitGroup{}
	errWg.Add(1) // TODO: we need to check if the return works when an error occurs and if it works if not
	go func(wg *sync.WaitGroup) {
		defer cancel()
		defer wg.Done()

		for err := range operationResponseChannel {
			// record the first error (the original error which should cause the other chunks to fail with canceled context)
			if err != nil && firstErr == nil {
				firstErr = err
				break    // Break the loop, call possible defer and end the routine
			}
		}
	}(&errWg)

	// Create the goroutines that process each operation (in parallel). Using
	// the wait groups to prevent race conditions when the channel is closed later
	wg := sync.WaitGroup{}
	for g := uint16(0); g < o.Concurrency; g++ {
		// One increment per routine because we want all operations to finish
		// before moving on when the listening channel closes
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			for f := range operationChannel {
				err := f()
				operationResponseChannel <- err
			}
		}(&wg)
	}

	// Add each chunk's operation to the channel.
	for chunkNum := uint16(0); chunkNum < numChunks; chunkNum++ {
		if firstErr != nil {
			// As soon as the first error occurs do not write any more jobs
			// because we want to return the first error of the func
			break 
		}

		curChunkSize := o.ChunkSize

		if chunkNum == numChunks-1 { // Last chunk
			curChunkSize = o.TransferSize - (int64(chunkNum) * o.ChunkSize) // Remove size of all transferred chunks from total
		}
		offset := int64(chunkNum) * o.ChunkSize

		operationChannel <- func() error {
			return o.Operation(offset, curChunkSize, ctx)
		}
	}
	close(operationChannel) // All operations were sent to the channel
	wg.Wait()               // Wait for all worker go routines to finish their work

	close(operationResponseChannel) // All sending go routines to the channel are done, the channel is now safe to close
	errWg.Wait()                    // Wait for the error worker to finish

	return firstErr
}
