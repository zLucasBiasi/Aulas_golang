package main

import "fmt"

func main() {

	arr := []int{10, 20, 30}

	if len(arr) == 0 {
		fmt.Println("o array esta vazio")
		return
	}
	number := 0

	for i := 0; i < len(arr); i++ {
		number += arr[i]
	}

	result := number / len(arr)

	fmt.Printf("a media é: %d \n", result)
}
