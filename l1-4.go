package main

import (
	"os/signal"
	"context"
	"syscall"
	"fmt"
	"time"
)

func exitSigTerm() {
	// Удобно, не надо создавать канал для отмены
	ctxCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, stop := signal.NotifyContext(ctxCancel, syscall.SIGINT)
	defer stop()

	channel := make(chan int)
	for i := 0; i < 5; i++ {
		idx := i
		go func() {
			select {
			case <-channel:
				fmt.Println("never")
			case <-ctx.Done():
				fmt.Printf("worker %d stopped\n", idx)
			}
		}()
	}

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		stop()
	}

	time.Sleep(1 * time.Second)
}