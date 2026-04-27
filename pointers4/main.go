package main

import "fmt"

func createPointer() *int {
	x := 23
	return &x
}

func main() {
	pX := createPointer()
	fmt.Println(*pX)

	*pX++

	fmt.Println(*pX)
}
