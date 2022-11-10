//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package shared

import (
	"context"
	"errors"
	"log"
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
	ops := make(chan func() error, o.Concurrency) // Create the channel that release 'concurrency' goroutines concurrently
	opErrors := make(chan error, o.Concurrency)   // Receives error responses from operations
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a background listener to the error channel. As soon as one error
	// is returned from an operation cancel all remaining operations and return
	// this first error.
	var firstErr error = nil

	// Create the goroutines that process each operation (in parallel). Using
	// the wait groups to prevent race conditions when the channel is closed later
	wg := sync.WaitGroup{}
	for g := uint16(0); g < o.Concurrency; g++ {
		// One increment per routine because we want all operations to finish
		// before moving on when the listening channel closes
		wg.Add(1)
		go func(wg *sync.WaitGroup, ops <-chan func() error, opResponses chan error, id uint16) {
			defer func() {
				log.Printf("closing worker %d", id)
				wg.Done()
			}()

			log.Printf("start worker %d", id)

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-opResponses:
					firstErr = err
					return
				case f, ok := <-ops:
					if !ok {
						return
					}
					err := f()
					if err != nil {
						log.Printf("sending err: %v", err)
						opResponses <- err
					}
				}
			}
		}(&wg, ops, opErrors, g)
	}

	// Add each chunk's operation to the channel.
	for chunkNum := uint16(0); chunkNum < numChunks; chunkNum++ {
		if firstErr != nil {
			// As soon as the first error occurs do not write any more jobs
			// because we want to return the first error of the func to the caller
			break
		}
		curChunkSize := o.ChunkSize

		if chunkNum == numChunks-1 { // Last chunk
			curChunkSize = o.TransferSize - (int64(chunkNum) * o.ChunkSize) // Remove size of all transferred chunks from total
		}
		offset := int64(chunkNum) * o.ChunkSize

		// This is a closure capturing the parameters and sending the operation to
		// the workers
		log.Printf("send to ops channel, chunkNum: %d out of %d chunks", chunkNum, numChunks)
		ops <- func() error {
			return o.Operation(offset, curChunkSize, ctx)
		}
	}

	log.Print("closing ops channel")
	close(ops) // All operations were sent to the channel
	wg.Wait()  // Wait gracefully for all worker go routines to finish their work

	log.Print("closing opResponses channel")
	close(opErrors) // All sending go routines to the channel are done, the channel is now safe to close

	return firstErr
}
