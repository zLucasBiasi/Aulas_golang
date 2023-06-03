package main

import "fmt"

type user struct {
	name  string
	idade uint8
}

func (u user) saveUser() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.name)
}

func (u user) Age() bool {
	return u.idade >= 18
}

func (u *user) birthday() {
	u.idade++
}

func main() {
	user1 := user{"Biasi", 20}
	fmt.Println(user1.idade)
	user1.saveUser()
	user1.birthday()
	fmt.Println(user1.idade)

	user2 := user{"Caixeta", 19}
	user2.saveUser()
	test := user2.Age()
	fmt.Println(test)
}
