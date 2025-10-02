package sorter

import(
	"context"
	"fmt"
)

const successResult = "file is sorted"

// checkSortedFunc проверка что файл отсортирован
func (s *Sorter) checkSortedFunc(ctx context.Context) {
	go func() {
		defer close(s.resultChan)

		var (
			prev string
			num  int64
		)

		for {
			select {
			case <-ctx.Done():
				return
			case line, ok := <-s.linesChan:
				if !ok {
					select {
					case <-ctx.Done():
					default:
						// Тут считаем что нигде не было ошибок
						// и канал был закрыт потому что прочитаны все строки
						s.resultChan <- successResult
					}
					return
				}

				num++
				if num == 1 {
					prev = line
					continue
				}

				if prev > line {
					s.resultChan <- fmt.Sprintf("disorder at line %d", num)
					return
				}

				prev = line
			}
		}
	}()
}
