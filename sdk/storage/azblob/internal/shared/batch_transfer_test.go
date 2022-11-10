package shared

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

var lock sync.Mutex
var sentinel atomic.Int64
var transferred int64

var transferSize int64 = int64(100_000)
var breakSize int64 = int64(20_000)

func TestDoBatchTransfer(t *testing.T) {
	// given a BatchTransferOptions
	// with an operation that mutates our sentinels
	bto := BatchTransferOptions{
		TransferSize:  transferSize,
		ChunkSize:     int64(10_000),
		Concurrency:   5,
		Operation:     successOperation,
		OperationName: "success-operation",
	}

	// when executing the batch-transfer
	DoBatchTransfer(context.Background(), &bto)

	// then our sentinels reflect the operation
	if transferred != transferSize {
		t.Logf("transferred: %d, wanted to transfer: %d", transferred, transferSize)
		t.Fail()
	}
}

func successOperation(offset int64, chunkSize int64, ctx context.Context) error {
	lock.Lock()
	transferred += chunkSize
	lock.Unlock()
	return nil
}

func TestDoBatchTransferError(t *testing.T) {
	// given a BatchTransferOptions
	// with an operation that mutates our sentinels
	bto := BatchTransferOptions{
		TransferSize:  transferSize,
		ChunkSize:     int64(5_000),
		Concurrency:   5,
		Operation:     errorOperation,
		OperationName: "error-operation",
	}

	// when executing the batch-transfer
	err := DoBatchTransfer(context.Background(), &bto)

	// we want to get an error
	if err == nil {
		t.Log("expected an error, but got none")
		t.FailNow()
	}

	if err.Error() != "transferred enough data" {
		t.Log("expected err: \"transferred enough data\", but got nil")
	}

	if transferred > transferSize+int64(10_000) {
		t.Logf("expected no more than %d, but got %d", breakSize, transferred)
		t.Fail()
	}

}

func errorOperation(offset int64, chunkSize int64, ctx context.Context) error {
	lock.Lock()
	defer lock.Unlock()

	log.Printf("processing offset: %d", offset)

	// /time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	select {
	case <-ctx.Done():
		log.Printf("worker received done signal: %d", offset)
		return errors.New("cancelled operation")
	default:
		transferred += chunkSize
		if transferred >= breakSize {
			log.Printf("transferred: %d, which is more than %d", transferred, breakSize)
			log.Print("throw error")
			return errors.New("transferred enough data")
		}
	}
	return nil
}
