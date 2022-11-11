package shared

import (
	"context"
	"errors"
	"log"
	"sync"
	"testing"
)

type transferred struct {
	sync.Mutex
	value int64
}

var tr transferred = transferred{}

var transferSize int64 = int64(10_000_000)
var breakSize int64 = int64(20_000)

func TestDoBatchTransfer(t *testing.T) {
	// given a BatchTransferOptions
	// with an operation that mutates our sentinels
	bto := BatchTransferOptions{
		TransferSize:  transferSize,
		ChunkSize:     int64(5_000),
		Concurrency:   5,
		Operation:     successOperation,
		OperationName: "success-operation",
	}

	// when executing the batch-transfer
	err := DoBatchTransfer(context.Background(), &bto)
	if err != nil {
		t.Logf("expected no error, but got: %v", err)
	}

	// then our sentinels reflect the operation
	if tr.value < transferSize {
		t.Logf("transferred: %d, wanted to transfer: %d", tr.value, transferSize)
		t.Fail()
	}
}

func successOperation(offset int64, chunkSize int64, ctx context.Context) error {
	tr.Lock()
	tr.value += chunkSize
	tr.Unlock()
	return nil
}

func TestDoBatchTransferError(t *testing.T) {
	// given a BatchTransferOptions
	// with an operation that mutates our sentinels
	chunkSize := int64(5_000)

	bto := BatchTransferOptions{
		TransferSize:  transferSize,
		ChunkSize:     chunkSize,
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
		t.Fail()
	}

}

func errorOperation(offset int64, chunkSize int64, ctx context.Context) error {
	//log.Printf("test: processing offset: %d", offset)

	select {
	case <-ctx.Done():
		log.Printf("test: context-cancelled: worker received done signal: %d", offset)
		return errors.New("cancelled operation")
	default:
		tr.Lock()
		defer tr.Unlock()
		tr.value += chunkSize
		//atomic.AddInt64(&transferred, chunkSize)

		if tr.value >= breakSize {
			//log.Printf("test: break-size exceeded: transferred: %d, which is more than %d", tr.value, breakSize)
			//log.Print("test: break-size exceeded: throw error")
			return errors.New("transferred enough data")
		}
	}
	return nil
}
