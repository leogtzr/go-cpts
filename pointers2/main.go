package main

import "fmt"

func main() {
	x := 10

	fmt.Println(x)
	update(&x, 11)

	fmt.Println(x)

	updateButNotReally(x, 12)
	fmt.Println(x)
}

func update(val *int, to int) {
	*val = to
}

func updateButNotReally(val, to int) {
	val = to
}
