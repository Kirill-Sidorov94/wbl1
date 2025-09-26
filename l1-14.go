package main

import (
	"fmt"
)

func typeDefenition(a any) {
	switch t := a.(type) {
	case int:
		fmt.Println("type - int")
	case string:
		fmt.Println("type - string")
	case bool:
		fmt.Println("type - bool")
	case chan any:
		fmt.Println("type - chan")
	default:
		fmt.Printf("unavailable type - %v", t)
	}
}