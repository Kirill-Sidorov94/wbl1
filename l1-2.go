package main

import (
	"fmt"
	"sync"
)

func CompetitiveSquaring() {
	arr := []int{2, 4, 8, 10}
	var wg sync.WaitGroup
	for i := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			fmt.Println(num * num)
		}(arr[i])
	}

	wg.Wait()
}