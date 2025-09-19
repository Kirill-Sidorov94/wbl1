package main

import (
	"fmt"
	"sync"
)

func concurrentEntryInMap() {	
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(2)
	go func () {
		defer wg.Done()
		sl := []int{1, 2, 3, 4 ,5 ,6}

		for i := range sl {
			m.Store(fmt.Sprintf("%d", sl[i]), sl[i])
		} 

	}()

	go func() {
		defer wg.Done()
		sl := []int{7 ,8 ,9, 10}

		for i := range sl {
			m.Store(fmt.Sprintf("%d", sl[i]), sl[i])
		} 
	}()

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
	   	fmt.Printf("%v: %v\n", key, value)
	    return true
	})
}