package main

import (
	"context"
	"time"
	"math/rand"
	"fmt"
)

func channelTimeout() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	channel := make(chan int, 2)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("sender stopped")
				return
			default:
				channel <- rand.Int()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("reader stopped")
				return
			case i, ok := <-channel:
				if !ok {
					fmt.Println("chanel closed or empty")
					return
				}
				fmt.Printf("result: %d\n", i)
			}
		}
	}()

	select {
	case <-ctx.Done():
		close(channel)
		cancel()
	}

	time.Sleep(1 * time.Second)
}