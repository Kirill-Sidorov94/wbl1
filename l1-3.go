package main

import(
	"context"
	"fmt"
	"sync"
	"time"
)

func MultipleWorkersRunning(workerCount int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case num, ok := <-ch:
					if !ok {
						return
					}
					fmt.Println(num)
				}
			}
		}()
	}

	for i := 0; i < 20000; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			ch <- i
		}
	}

	close(ch)
	wg.Wait()
}