package output

import (
	"context"
	"fmt"
)

type Output struct {
	resultChan <-chan string
	errChan chan<- error
	doneChan chan struct{}
}

func New(resultCh <-chan string, errChan chan<- error) *Output {
	return &Output{
		resultChan: resultCh,
		errChan: errChan,
		doneChan: make(chan struct{}),
	}
}

func (o *Output) Write(ctx context.Context) {
	go func() {
		defer close(o.doneChan) 

		for {
			select {
			case <-ctx.Done():
				return
			case r, ok := <-o.resultChan:
				if !ok {
					return
				}

				fmt.Println(r)
			}
		}
	}()
}

func (o *Output) Done() <-chan struct{} {
	return o.doneChan
}