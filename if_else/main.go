package main

import "fmt"

func main() {

	numberOne := 10

	if numberOne > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if numberTwo := 10; numberTwo == numberOne {
		fmt.Println("numeros iguais")
	} else {
		fmt.Println("numeros diferentes")
	}

	// fmt.Println(numberTwo)
	// eu nao poderia chamar essa variavel fora porque o if init respeita apenas o escopo
}
