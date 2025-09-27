package main

import (
	"fmt"
)

func intersectionOfSets() {
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{2, 3, 4, 5, 6, 7, 8}
	var size int

	if len(sl1) > len(sl2) {
		size = len(sl1)
	} else {
		size = len(sl2)
	}

	m := make(map[int]struct{}, size)
	result := make([]int, 0, size)

	for i := range sl1 {
		m[sl1[i]] = struct{}{}
	}

	for i := range sl2 {
		if _, ok := m[sl2[i]]; ok {
			result = append(result, sl2[i])
		}
	}

	fmt.Println(result)
}