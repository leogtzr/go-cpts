package main

import "fmt"

func main() {
	var a int = 25
	var b *int = &a
	var c **int = &b

	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(**c)
}
