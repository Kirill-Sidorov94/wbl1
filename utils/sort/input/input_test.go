package input

import (
	"bufio"
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestInput_Read_Stdin(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdin = r

	testData := "line1\nline2\nline3\n"
	go func() {
		defer w.Close()
		w.WriteString(testData)
	}()

	errChan := make(chan error, 1)
	input := New("", errChan)

	ctx := context.Background()
	input.Read(ctx)

	var lines []string
	for line := range input.GetLinesChan() {
		lines = append(lines, line)
	}

	select {
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	default:
	}

	expectedLines := []string{"line1", "line2", "line3"}
	if len(lines) != len(expectedLines) {
		t.Errorf("expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, expected := range expectedLines {
		if lines[i] != expected {
			t.Errorf("line %d: expected '%s', got '%s'", i, expected, lines[i])
		}
	}
}

func TestInput_Read_File(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	testData := "file_line1\nfile_line2\nfile_line3\n"
	if _, err := tmpfile.WriteString(testData); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()

	errChan := make(chan error, 1)
	input := New(tmpfile.Name(), errChan)

	ctx := context.Background()
	input.Read(ctx)

	var lines []string
	for line := range input.GetLinesChan() {
		lines = append(lines, line)
	}

	select {
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	default:
	}

	expectedLines := []string{"file_line1", "file_line2", "file_line3"}
	if len(lines) != len(expectedLines) {
		t.Errorf("expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, expected := range expectedLines {
		if lines[i] != expected {
			t.Errorf("line %d: expected '%s', got '%s'", i, expected, lines[i])
		}
	}
}

func TestInput_Read_FileNotFound(t *testing.T) {
	errChan := make(chan error, 1)
	input := New("nonexistent_file.txt", errChan)

	ctx := context.Background()
	input.Read(ctx)

	time.Sleep(100 * time.Millisecond)

	select {
	case _, ok := <-input.GetLinesChan():
		if ok {
			t.Error("expected linesChan to be closed")
		}
	default:
		t.Error("linesChan should be closed when file doesn't exist")
	}

	select {
	case err := <-errChan:
		if err == nil {
			t.Error("expected error for nonexistent file")
		}
	default:
		t.Error("expected error to be sent to errChan")
	}
}

func TestInput_Read_CancelledContext(t *testing.T) {
	errChan := make(chan error, 1)
	input := New("", errChan)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Отменяем контекст сразу

	input.Read(ctx)

	select {
	case _, ok := <-input.GetLinesChan():
		if ok {
			t.Error("expected linesChan to be closed")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("linesChan should be closed immediately with cancelled context")
	}

	select {
	case err := <-errChan:
		t.Errorf("unexpected error: %v", err)
	default:
	}
}

func TestReadLines_ScannerError(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("test"))
	
	linesChan := make(chan string, 10)
	ctx := context.Background()
	
	err := readLines(ctx, scanner, linesChan)
	if err != nil {
		t.Errorf("unexpected error with valid scanner: %v", err)
	}
}

// createTestFile вспомогательная функция для создания тестового файла
func createTestFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "test_input")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()
	
	return tmpfile.Name()
}