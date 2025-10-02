package output

import (
	"context"
	"testing"
	"time"
	"fmt"
)

func TestOutput_Write_Success(t *testing.T) {
	ctx := context.Background()
	resultChan := make(chan string, 1000)
	errChan := make(chan error, 1)

	resultChan <- "line1"
	resultChan <- "line2" 
	resultChan <- "line3"
	close(resultChan)

	output := New(resultChan, errChan)
	output.Write(ctx)

	select {
	case <-output.Done():
	case <-time.After(100 * time.Millisecond):
		t.Error("output didn't complete in time")
	}
}

func TestOutput_Write_Empty(t *testing.T) {
	ctx := context.Background()
	resultChan := make(chan string, 1000)
	errChan := make(chan error, 1)

	close(resultChan)

	output := New(resultChan, errChan)
	output.Write(ctx)

	select {
	case <-output.Done():
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("output didn't complete with empty input")
	}
}

func TestOutput_Write_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan string, 1000)
	errChan := make(chan error, 1)

	output := New(resultChan, errChan)
	output.Write(ctx)

	resultChan <- "test"
	
	time.Sleep(10 * time.Millisecond)
	cancel()

	select {
	case <-output.Done():
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("output didn't respect context cancellation")
	}
}

func TestOutput_Write_WithError(t *testing.T) {
	ctx := context.Background()
	resultChan := make(chan string, 2)
	errChan := make(chan error, 1)

	errChan <- fmt.Errorf("test error")
	resultChan <- "line1"
	resultChan <- "line2"
	close(resultChan)

	output := New(resultChan, errChan)
	output.Write(ctx)

	select {
	case <-output.Done():
		// Успех
	case <-time.After(100 * time.Millisecond):
		t.Error("output didn't complete with error in channel")
	}
}

func TestOutput_Done_ChannelClosed(t *testing.T) {
	ctx := context.Background()
	resultChan := make(chan string)
	errChan := make(chan error, 1)

	output := New(resultChan, errChan)
	output.Write(ctx)

	close(resultChan)

	select {
	case <-output.Done():
		// Успех - канал закрыт
	case <-time.After(100 * time.Millisecond):
		t.Error("Done() channel wasn't closed")
	}
}