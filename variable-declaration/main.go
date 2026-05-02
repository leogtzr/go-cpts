package main

import "fmt"

func main() {
	var (
		b int
		c string
		d float64
	)

	b = 1
	c = "Hello World"
	d = 0.344

	fmt.Println(b, c, d)

	// the following is invalid, since there is no init and no type, we do not know what it is
	// var a
}
