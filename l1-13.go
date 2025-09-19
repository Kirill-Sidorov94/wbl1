package main

import (
	"fmt"
)

func exchangeOfValues() {
	a := 10
	b := 15
	
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)
}