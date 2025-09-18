package main

// атомик используют только на низкоуровневых операциях и не рекомендуют
// https://pkg.go.dev/sync/atomic
// мьютекс же сведет на нет паралельность если это просто счетчик
import (
	"context"
	"time"
	"fmt"
	"sync"
)

type Counter struct {
	channel chan int
	value   int
}

func newCounter() *Counter {
	return &Counter{
		channel: make(chan int),
	}
}

func (c *Counter) run() {
	go func() {
		for {
			select {
			case _, ok := <-c.channel:
				if !ok {
					return
				}

				c.value++
			}
		}
	}()
}

func (c *Counter) inc() bool {
	select {
	case c.channel <- 1:
		return true
	default:
		return false
	}
}

func (c *Counter) stop() {
	close(c.channel)
}

func (c *Counter) getValue() int {
	return c.value
}

func competitiveCounter(workerCount int) {
	ctx, cancel := context.WithTimeout(context.Background(), 4 * time.Second)
	defer cancel()

	counter := newCounter()
	counter.run()

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					counter.inc()
				}
			}
		}()
	}

	wg.Wait()
	counter.stop()
	fmt.Println(counter.getValue())
}