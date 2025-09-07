package main

import(
	"fmt"
)

func main() {
	// l1.1 go run l1-1.go main.go
	a := Action{Human: Human{Name: "action"}}
	fmt.Println(a.GetName())
}