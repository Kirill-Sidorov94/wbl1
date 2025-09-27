package main

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

var (
	numbersOnlyRegex = regexp.MustCompile(`^[0-9]+$`)
	pattern = regexp.MustCompile(`(\p{L}|\\\d)(\d*)`)
)

func unpackingString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	
	if numbersOnlyRegex.MatchString(str) {
		return "", fmt.Errorf("str - %s contains only multipliers", str)
	}

	return pattern.ReplaceAllStringFunc(str, func(match string) string {
		parts := pattern.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}

		subStr := parts[1]
		multiplier, err := strconv.Atoi(parts[2]);
		if err != nil {
			multiplier = 1
		}

		runes := []rune(subStr)
		if string(runes[0]) == "\\" {
			return strings.Repeat(string(runes[1]), multiplier)
		} else {
			return strings.Repeat(string(runes[0]), multiplier)
		}
    }), nil
}