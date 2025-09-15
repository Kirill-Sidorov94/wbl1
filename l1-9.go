package main

import (
	"fmt"
	"sync"
	"context"
	"math/rand"
	"time"
)

func conveyorOfNumbers() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())

	var arr [20000]int
    for i := 0; i < len(arr); i++ {
        arr[i] = rand.Int()
    }

	channel1 := make(chan int, 5)
	channel2 := make(chan int, 5)
	var (
		counter       int
		handleCounter int
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case num, ok := <-channel1:
				if !ok {
					return
				}

				select{
				case <-ctx.Done():
					return
				default:
					channel2 <- num
					counter++

					if counter == len(arr) {
						close(channel2)
					}
				}

			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case num, ok := <-channel2:
				if !ok {
					return
				}

				fmt.Println(num*2)
				handleCounter++
			}
		}
	}()

	for i := 0; i < len(arr); i++ {
		select {
		case <-ctx.Done():
			break
		default:
			channel1 <- arr[i]
		}
	}
	close(channel1)
	wg.Wait()

	fmt.Printf("programm close, handle %d item\n", handleCounter)
}