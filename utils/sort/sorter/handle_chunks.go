package sorter

import (
	"context"
	"bufio"
	"os"
)

// mergeWithHeap мержим файлы 
func (s *Sorter) mergeWithHeap(ctx context.Context, tmpChunkFiles []string) {
	defer close(s.resultChan)

	scanners := make([]*bufio.Scanner, 0, len(tmpChunkFiles))
	tmpFiles := make([]*os.File, 0, len(tmpChunkFiles))
	
	for i := range tmpChunkFiles {
		file, err := os.Open(tmpChunkFiles[i])
		if err != nil {
			s.errChan <- err
			return
		}
		tmpFiles = append(tmpFiles, file)
		scanners = append(scanners, bufio.NewScanner(file))
	}

	mergeHeap := newMergeHeap(scanners, s.less)

	var prev string
	for mergeHeap.Len() > 0 {
		select {
		case <-ctx.Done():
			return
		default:
			line, err := mergeHeap.next()
			if err != nil {
				s.errChan <- err
				return
			}
			
			// Применяем unique
			if s.unique && line == prev {
				continue
			}
			
			s.resultChan <- line
			prev = line
		}
	}

	for _, file := range tmpFiles {
		file.Close()
	}
}
