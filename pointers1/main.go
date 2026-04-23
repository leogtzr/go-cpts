package main

import "fmt"

type Book struct {
	isbn  int
	title string
}

func main() {
	var a int = 10
	var b *int

	b = &a

	println(b)
	println(*b)

	*b = 82

	fmt.Println(*b)
	fmt.Println(a)

	fmt.Println("--- addresses ---")
	fmt.Println(&a)
	fmt.Println(&b)

	// Check the address in decimal, instead of HEX....
	fmt.Printf("Val: %d\n", b)

	var p1 *int
	p2 := new(int)

	fmt.Println(p1)
	fmt.Println(p2) // points somewhere ...
	fmt.Println(*p2)

	*p2 = 351

	fmt.Println(*p2)

	p1 = new(int)
	*p1 = 3334

	fmt.Println(*p1)

	val := &Book{}
	val.isbn = 453
	val.title = "Stoner"

	val2 := new(Book)
	val2.isbn = 234
	val2.title = "Conde de Montecristo"

	fmt.Println(val)
	fmt.Println(val2)
	fmt.Println(val2.title)

	// Pointers are type safe ...
	// Cannot do the following:
	// var aa int = 10
	// var bb *string = &aa
}
