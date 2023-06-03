package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	total := soma(1, 2, 3, 4)
	fmt.Println(total)
}
