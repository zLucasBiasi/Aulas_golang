package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u *user) changeName(name string) {
	u.name = name
}

func main() {
	var u user

	u.name = "Biasi"

	fmt.Println(u.name)
	u.changeName("Caixeta")
	fmt.Println(u.name)
}
