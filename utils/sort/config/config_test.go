package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"numeric sort", []string{"-n", "file.txt"}, false},
		{"reverse numeric", []string{"-nr", "file.txt"}, false},
		{"key column sort", []string{"-k", "2", "file.txt"}, false},
		{"multiple flags", []string{"-nru", "file.txt"}, false},
		{"mixed order", []string{"file.txt", "-n", "file1.txt"}, false},
		{"error -k without number", []string{"-k"}, true},
		{"error -k with text", []string{"-k", "abc"}, true},
		{"error -k with zero", []string{"-k", "0"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = append([]string{"mysort"}, tt.args...)

			_, err := New()

			if tt.wantErr && err == nil {
				t.Error("Expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}