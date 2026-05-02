package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("vim-go")

	g := 10
	h, i := "hello", 20

	fmt.Println(g, h, i)

	// The following is valid, since "j" is new
	// as long one is new, we call this re-declaration and it is OK for h

	h, j := "world", 7

	fmt.Println(h, j)

	name := "Leo"
	name, age := "Perla", 34

	fmt.Println(name)
	fmt.Println(age)

	leo := Person{
		Name: "Leo",
		Age:  35,
	}

	fmt.Println(leo)

	// The following fails because even that there is a new variable ("name"), the left variable has to be a plain variable
	// leo.Age, name := 36, "juanis"
}
