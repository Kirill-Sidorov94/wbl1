package sorter

import (
	"context"
	"sync"
)

// applySort применяем сортировку
func (s *Sorter) applySort(ctx context.Context) {
	tmpFileNamesChan := make(chan string)
	var tempFileNames []string
	go func () {
		defer func() {
			close(tmpFileNamesChan)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case tmpfileName, ok := <-tmpFileNamesChan:
				if !ok {
					return
				}
				tempFileNames = append(tempFileNames, tmpfileName)
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < s.workerCount; i++ {
		idx := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			chunk := s.chunkPool.Get().([]string)
			chunk = chunk[:0]

			for {
				select {
				case <-ctx.Done():
					return
				case line, ok := <- s.linesChan:
					if !ok {
						if len(chunk) > 0 {
							s.createSortedChunk(ctx, chunk, idx, tmpFileNamesChan)
						}
						return
					}

					chunk = append(chunk, line)
					if len(chunk) >= s.chunkSize {
						s.createSortedChunk(ctx, chunk, idx, tmpFileNamesChan)
						chunk = s.chunkPool.Get().([]string)
						chunk = chunk[:0]
					}
				}
			}
		}()
	}
	wg.Wait()

	s.mergeWithHeap(ctx, tempFileNames)
}
