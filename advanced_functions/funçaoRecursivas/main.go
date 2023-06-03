package main

import "fmt"

func fibonacci(posiçao uint) uint {
	if posiçao <= 1 {
		return posiçao
	}

	return fibonacci(posiçao-2) + fibonacci(posiçao-1)
}

func main() {

	posiçao := uint(10)

	fmt.Println(fibonacci(posiçao))
}
