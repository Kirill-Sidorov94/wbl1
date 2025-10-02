package input

import (
	"os"
	"bufio"
	"context"
)

const (
	linesCount = 5000
	stdinKind = "stdin"
	fileKind  = "file"
)

// Input инпут
type Input struct {
	file string
	kind string
	linesChan chan string
	errChan chan<- error
}

// New реализует инпут с настроками чтения
func New(file string, errChan chan<- error) *Input {
	kind := stdinKind
	if file != "" {
		kind = fileKind
	}

	return &Input{
		file: file, 
		kind: kind, 
		linesChan: make(chan string, linesCount), 
		errChan: errChan,
	}
}

// Read запускает чтение из источника
func (i *Input) Read(ctx context.Context) {
	go func() {
		defer close(i.linesChan)
		var err error

		select {
		case <-ctx.Done():
			return
		default:
			if i.kind == stdinKind {
				err = i.readStdin(ctx)
			} else {
				err = i.readFile(ctx)
			}
		}

		select {
		case <-ctx.Done():
			return
		default:
			if err != nil {
				i.errChan <- err
			}
		}
	}()
}

// GetLinesChan возвращает канал с прочитаными строками
func (i *Input) GetLinesChan() <-chan string {
	return i.linesChan
}

// readFile читает из файла
func (i *Input) readFile(ctx context.Context) error {
	file, err := os.Open(i.file)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := readLines(ctx, scanner, i.linesChan); err != nil {
		return err
	}

	return nil
}

// readStdin читает из буфера
func (i *Input) readStdin(ctx context.Context) error {
	scanner := bufio.NewScanner(os.Stdin)
	if err := readLines(ctx, scanner, i.linesChan); err != nil {
		return err
	}

	return nil
}

// readLines процесс чтения строки
func readLines(ctx context.Context, scanner *bufio.Scanner, linesChan chan<- string) error {
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil
		default:
			linesChan <- scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}