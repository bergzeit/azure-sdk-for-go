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

type firstError struct {
	*sync.Mutex
	error error
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
	ops := make(chan func() error, o.Concurrency) // Create the channel that release 'concurrency' goroutines concurrently
	opErrors := make(chan error)                  // Receives error responses from operations
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// an instance for the first error ever occured in a worker
	firstErr := firstError{
		&sync.Mutex{},
		nil,
	}

	// Create the goroutines that process each operation (in parallel). Using
	// the wait groups to prevent race conditions when the channel is closed later
	// Worker routines are done, when the ops channel is closed.
	wg := sync.WaitGroup{}
	for g := uint16(0); g < o.Concurrency; g++ {
		// One increment per routine because we want all operations to finish
		// before moving on when the listening channel closes
		wg.Add(1)
		go func(wg *sync.WaitGroup, ops <-chan func() error, opErrs chan<- error) {
			defer wg.Done()

			for f := range ops {
				err := f()
				if err != nil {
					opErrs <- err
				}
			}
		}(&wg, ops, opErrors)
	}

	// Add each chunk's operation to the channel.
	wg.Add(1)
	go func(ctx context.Context, op chan func() error, wg *sync.WaitGroup) {
		defer func() {
			close(op) // All operations were sent to the channel, so the workers are done
			wg.Done()
		}()

		for chunkNum := uint16(0); chunkNum < numChunks; chunkNum++ {
			select {
			case <-ctx.Done():
				return
			default:
				curChunkSize := o.ChunkSize

				if chunkNum == numChunks-1 { // Last chunk
					curChunkSize = o.TransferSize - (int64(chunkNum) * o.ChunkSize) // Remove size of all transferred chunks from total
				}
				offset := int64(chunkNum) * o.ChunkSize

				// This is a closure capturing the parameters and sending the operation to
				// the workers
				ops <- func() error {
					return o.Operation(offset, curChunkSize, ctx)
				}
			}
		}
	}(ctx, ops, &wg)

	// reading the errors returned from the workers, so that we can get
	// the first error thrown and save it as return value
	var errWg sync.WaitGroup
	errWg.Add(1)
	go func(opErrs <-chan error, fe *firstError, wg *sync.WaitGroup) {
		defer wg.Done()
		for err := range opErrs {
			// hard locking the error
			fe.Lock()
			// if we've catched an error already, we continue
			// thus ensure, the error channel is emptied
			if fe.error != nil {
				fe.Unlock()
				continue
			}
			fe.error = err
			fe.Unlock()
			cancel()
		}
	}(opErrors, &firstErr, &errWg)

	wg.Wait() // Wait gracefully for all worker go routines to finish their work

	close(opErrors) // All sending go routines to the channel are done, the channel is now safe to close

	errWg.Wait()

	return firstErr.error
}
