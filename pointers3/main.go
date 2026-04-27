package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u User) updateName(name string) {
	u.Name = name
}

// The method is receiving a pointer to user.
func (u *User) changeName(name string) {
	u.Name = name
}

func main() {
	user := User{ID: 35, Name: "Leonardo"}
	fmt.Println(user)

	user.updateName("León")
	fmt.Println(user)

	user.changeName("Tobias")
	fmt.Println(user)
}
