package main

import "fmt"

func closure() func() {
	texto := "Dentro da funçao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da func main"
	fmt.Println(texto)

	funcnova := closure()
	funcnova()
}
