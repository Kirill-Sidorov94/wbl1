package sorter

import (
	"context"
	"os"
	"testing"
	"time"
	"strings"

	"github.com/Kirill-Sidorov94/wbl1/utils/sort/config"
)

func TestSorter_WithTestFile(t *testing.T) {
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		t.Skip("test.txt not found, skipping file-based tests")
	}

	tests := []struct {
		name     string
		cfg      *config.Config
		checkFn  func([]string) bool
	}{
		{
			name: "default sort",
			cfg:  &config.Config{},
			checkFn: func(lines []string) bool {
				for i := 1; i < len(lines); i++ {
					if lines[i] < lines[i-1] {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			content, err := os.ReadFile("test.txt")
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}
			lines := strings.Split(strings.TrimSpace(string(content)), "\n")
			linesCh := make(chan string, len(lines))
			errCh := make(chan error, 1)
			for i := range lines {
				linesCh <- lines[i]
			}

			sorter := New(&Options{
		        LinesChan: linesCh, 
		        ErrChan: errCh, 
		        CheckSorted: tt.cfg.CheckSorted,
		        Reverse: tt.cfg.Reverse,
		        Unique: tt.cfg.Unique,
		        SortType: tt.cfg.SortType,
		        KeyColumn: tt.cfg.KeyColumn,
		        IgnoreBlanks: tt.cfg.IgnoreBlanks,
		    })

			sorter.Process(ctx)

			var sorted []string
			for l := range sorter.GetResultChan() {
				sorted = append(sorted, l)
			}

			if !tt.checkFn(sorted) {
				t.Errorf("Sort check failed for %s", tt.name)
			}

			t.Logf("%s: sorted %d lines", tt.name, len(sorted))
		})
	}
}
