package main

import (
	"fmt"
)

func exchangeOfValues() {
	a := 10
	b := 15
	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}