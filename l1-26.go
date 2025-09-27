package main

import (
	"strings"
)

func uniqSymbolsToStr(str string) bool {
	lowerStr := strings.ToLower(str)
	runes := []rune(lowerStr)
	m := make(map[rune]struct{}, len(runes))

	for i := range runes {
		if _, ok := m[runes[i]]; ok {
			return false
		}
		m[runes[i]] = struct{}{}
	}

	return true
}