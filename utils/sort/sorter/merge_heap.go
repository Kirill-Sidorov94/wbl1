package sorter

import (
	"bufio"
	"container/heap"
	"io"
)

// HeapItem представляет элемент в heap'е
type HeapItem struct {
	line   string    // текущая строка
	reader *bufio.Scanner // сканер чанка
}

// MergeHeap реализует heap.Interface для слияния чанков 
// https://pkg.go.dev/container/heap
type MergeHeap struct {
	items []*HeapItem
	less  func(a, b string) bool
}

func (h MergeHeap) Len() int { 
	return len(h.items) 
}

func (h MergeHeap) Less(i, j int) bool {
	return h.less(h.items[i].line, h.items[j].line)
}

func (h MergeHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MergeHeap) Push(x interface{}) {
	h.items = append(h.items, x.(*HeapItem))
}

func (h *MergeHeap) Pop() interface{} {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[0 : n-1]
	return item
}

// newMergeHeap создает и инициализирует heap для слияния
func newMergeHeap(scanners []*bufio.Scanner, less func(a, b string) bool) *MergeHeap {
	h := &MergeHeap{less: less}
	
	for _, scanner := range scanners {
		if scanner.Scan() {
			item := &HeapItem{
				line:   scanner.Text(),
				reader: scanner,
			}
			h.items = append(h.items, item)
		}
	}
	
	heap.Init(h)
	return h
}

// next возвращает следующую минимальную строку
func (h *MergeHeap) next() (string, error) {
	if len(h.items) == 0 {
		return "", io.EOF
	}

	// Берем минимальный элемент
	item := h.items[0]
	minLine := item.line

	// Читаем следующую строку из этого сканера
	if item.reader.Scan() {
		item.line = item.reader.Text()
		heap.Fix(h, 0)
	} else {
		// Сканер закончился - удаляем из heap
		heap.Pop(h)
	}

	return minLine, nil
}
