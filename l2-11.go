package main

import (
	"strings"
	"sort"
)

func groupAnagrams(strings []string) map[string][]string {
	if len(strings) == 0 {
		return make(map[string][]string)
	}

	result := make(map[string][]string, len(strings))
	anagramsMap := make(map[string]string, len(strings))

	for i := range strings {
		str := strings.ToLower(strings[i])
		runes := []rune(str)
		sort.Slice(runes, func(a, b int) bool {
			return runes[a] < runes[b]
		})
		sortStr := string(runes)
		value, ok := anagramsMap[sortStr]
		if ok {
			result[value] = append(result[value], str)
		} else {
			anagramsMap[sortStr] = str
			result[str] = make([]string, 0, len(strings))
			result[str] = append(result[str], str)
		}
	}

	for key, group := range result {
		if len(group) <= 1 {
			delete(result, key)
		}
	}

	return result
}