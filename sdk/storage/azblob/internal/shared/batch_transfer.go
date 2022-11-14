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

type firstError struct {
	lock  *sync.Mutex
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
	opErrors := make(chan error, o.Concurrency)   // Receives error responses from operations
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// an instance for the first error ever occured in a worker
	firstErr := firstError{
		lock: &sync.Mutex{},
	}

	// Create the goroutines that process each operation (in parallel). Using
	// the wait groups to prevent race conditions when the channel is closed later
	// Worker routines are done, when the ops channel is closed.
	wg := sync.WaitGroup{}
	for g := uint16(0); g < o.Concurrency; g++ {
		// One increment per routine because we want all operations to finish
		// before moving on when the listening channel closes
		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup, ops <-chan func() error, opErrs chan<- error, id uint16, fe *firstError) {
			defer func() {
				log.Printf("worker #%d: closing", id)
				wg.Done()
			}()

			log.Printf("worker #%d: start", id)

			for f := range ops {
				err := f()
				if err != nil {
					log.Printf("worker #%d: sending err: %v", id, err)
					opErrs <- err
				} else {
					log.Printf("worker #%d: success", id)
				}
			}
		}(ctx, &wg, ops, opErrors, g, &firstErr)
	}

	// Add each chunk's operation to the channel.
	var mainWg sync.WaitGroup
	mainWg.Add(1)
	go func(ctx context.Context, op chan func() error, wg *sync.WaitGroup) {
		defer func() {
			close(op) // All operations were sent to the channel, so the workers are done
			wg.Done()
		}()

		for chunkNum := uint16(0); chunkNum < numChunks; chunkNum++ {
			select {
			case <-ctx.Done():
				log.Print("batch-main: stop sending jobs, due to context-cancel")
				return
			default:
				curChunkSize := o.ChunkSize

				if chunkNum == numChunks-1 { // Last chunk
					curChunkSize = o.TransferSize - (int64(chunkNum) * o.ChunkSize) // Remove size of all transferred chunks from total
				}
				offset := int64(chunkNum) * o.ChunkSize

				// This is a closure capturing the parameters and sending the operation to
				// the workers
				log.Printf("batch-main: send to ops channel, chunkNum: %d out of %d chunks", chunkNum, numChunks)
				ops <- func() error {
					return o.Operation(offset, curChunkSize, ctx)
				}
			}
		}
	}(ctx, ops, &mainWg)

	// reading the errors returned from the workers, so that we can get
	// the first error thrown and save it as return value
	var errWg sync.WaitGroup
	errWg.Add(1)
	go func(opErrs <-chan error, fe *firstError, wg *sync.WaitGroup) {
		defer wg.Done()
		for err := range opErrs {
			// hard locking the error
			fe.lock.Lock()
			// if we've catched an error already, we continue
			// thus ensure, the error channel is emptied
			if fe.error != nil {
				fe.lock.Unlock()
				continue
			}
			log.Printf("batch-main: received error: %v", err)
			fe.error = err
			fe.lock.Unlock()
			cancel()
		}
	}(opErrors, &firstErr, &errWg)

	mainWg.Wait()
	log.Print("batch-main: closing ops channel")
	wg.Wait() // Wait gracefully for all worker go routines to finish their work

	log.Print("batch-main: closing opResponses channel")
	close(opErrors) // All sending go routines to the channel are done, the channel is now safe to close

	errWg.Wait()

	return firstErr.error
}
