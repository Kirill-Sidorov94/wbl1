package main

import (
	"testing"
)

func TestUnpackingString_Comprehensive(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        // Basic functionality
        {"basic", "a4bc2d5e", "aaaabccddddde", false},
        {"no multipliers", "abcd", "abcd", false},
        {"single char", "a4", "aaaa", false},
        {"all ones", "a1b1c1", "abc", false},
        
        // Escape sequences
        {"escape digit", "a2\\3b1", "aa3b", false},
        {"multiple escapes", "\\4\\5\\6", "456", false},
        {"escape with multiplier", "\\45", "44444", false},
        {"mixed escapes", "a2\\3b1\\42", "aa3b44", false},
        
        // Edge cases
        {"empty", "", "", false},
        {"zero multiplier", "a0", "", false},
        {"leading zero", "a01", "a", false},
        {"large multiplier", "a10", "aaaaaaaaaa", false},
        {"only backslash", "\\\\", "\\\\", false},
        
        // Complex cases
        {"complex 1", "a4b2c3d1e0f5", "aaaabbcccdfffff", false},
        {"complex with escapes", "a2\\3b1\\\\2c4", "aa3b\\2cccc", false},
        {"multiple digits escape", "\\410", "4444444444", false},
        {"mixed case", "A2b3C1", "AAbbbC", false},
        
        // Error cases
        {"only numbers", "12345", "", true},
        {"only multipliers", "123", "", true},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := unpackingString(tc.input)
            
            if tc.wantErr {
                if err == nil {
                    t.Errorf("expected error, got nil")
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
                if result != tc.expected {
                    t.Errorf("got %q, want %q", result, tc.expected)
                }
            }
        })
    }
}