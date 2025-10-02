package sorter

import (
	"context"
	"runtime"
	"sync"
	"os"
)

// Options опции сортера
type Options struct {
	LinesChan <-chan string
	ErrChan chan<- error
	WorkerCount int
	ChunkSize   int64
	CheckSorted bool
	Reverse     bool
	Unique      bool
	SortType    string
	KeyColumn   int
	IgnoreBlanks bool
}

// Sorter структура хранящая все для сортировки
type Sorter struct {
	linesChan <-chan string
	errChan chan<- error
	resultChan chan string
	workerCount int
	chunkSize   int
	chunkPool sync.Pool
	tempDir   string
	checkSorted bool
	reverse     bool
	unique      bool
	sortType string
	keyColumn int
	ignoreBlanks bool
}

// New инициализирует Sorter с настроками
func New(o *Options) *Sorter {
	s := &Sorter{
		linesChan: o.LinesChan,
		errChan: o.ErrChan,
		resultChan: make(chan string, 1000),
		workerCount: o.WorkerCount,
		tempDir: os.TempDir(),
		checkSorted: o.CheckSorted,
		reverse: o.Reverse,
		unique: o.Unique,
		sortType: o.SortType,
		keyColumn: o.KeyColumn,
		ignoreBlanks: o.IgnoreBlanks,
	}

	// Кол-во воркеров которые будут обрабатывать чанки
	if s.workerCount == 0 {
		s.workerCount = runtime.NumCPU()
	}

	// Пул для экономии
	if o.ChunkSize == 0 {
		s.chunkSize = 5000
	}
	s.chunkPool = sync.Pool{
		New: func() interface{} {
			return make([]string, 0, s.chunkSize) 
		},
	}

	return s
}

// Process основной процесс
func (s *Sorter) Process(ctx context.Context) {
	if s.checkSorted {
		s.checkSortedFunc(ctx)
		return
	}

	s.applySort(ctx)
}

// GetResultChan возвращает канал с отсортированными данными
func (s *Sorter) GetResultChan() <-chan string {
	return s.resultChan
}
