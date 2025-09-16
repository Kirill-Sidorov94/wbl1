package main

import (
	"fmt"
)

func setOfStrings() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool, len(s))
	r := make([]string, 0, len(s))

	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = true
			r = append(r, s[i])
		}
	}

	fmt.Println(r)
}