package main

import (
	"fmt"
	"sync"
	"context"
	"time"
	"runtime"
)

func stopGoRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("worker 1 stopped because code done")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var counter int
		for {
			if counter == 10 {
				fmt.Println("worker 2 stopped because condition done")
				return
			}

			counter++
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("worker 3 stopped because context done")
			return
		}
	}()

	channel := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case i, _ := <-channel:
				fmt.Printf("worker 4 stopped because read channel signal: %d\n", i)
				return 
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("worker 5 stopped because call Goexit")
		runtime.Goexit()
	}()

	time.Sleep(1 * time.Second)
	channel <- 0

	wg.Wait()
}