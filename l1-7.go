package main

import (
	"fmt"
	"sync"
	"github.com/orcaman/concurrent-map/v2"
)

func concurrentEntryInMap() {	
	m := cmap.New[int]()

	var wg sync.WaitGroup
	wg.Add(2)
	go func () {
		defer wg.Done()
		sl := []int{1, 2, 3, 4 ,5 ,6}

		for i := range sl {
			m.Set(fmt.Sprintf("%d", sl[i]), sl[i])
		} 

	}()

	go func() {
		defer wg.Done()
		sl := []int{7 ,8 ,9, 10}

		for i := range sl {
			m.Set(fmt.Sprintf("%d", sl[i]), sl[i])
		} 
	}()

	wg.Wait()
}