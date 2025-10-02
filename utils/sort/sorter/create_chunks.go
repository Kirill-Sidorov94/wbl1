package sorter

import (
	"context"
	"sort"
	"path/filepath"
	"strconv"
	"bufio"
	"os"
)

// createSortedChunk создаем отсортированные временые файлы
func (s *Sorter) createSortedChunk(ctx context.Context, chunk []string, chunkIdx int, tmpFilesCh chan<- string) {
	sort.Slice(chunk, func(i, j int) bool {
		return s.less(chunk[i], chunk[j])
	})

	filename := filepath.Join(s.tempDir, "sort_chunk_"+strconv.Itoa(chunkIdx)+".tmp")
	tmpFile, err := os.Create(filename)
	if err != nil {
		s.errChan <- err
		return
	}
	defer tmpFile.Close()

	writer := bufio.NewWriter(tmpFile)
	for i := range chunk {
		writer.WriteString(chunk[i] + "\n")
	}
	writer.Flush()

	select {
	case <-ctx.Done():
	default:
		tmpFilesCh <- tmpFile.Name()
	}
}
