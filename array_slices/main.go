package main

import "fmt"

func main() {

	var arr [5]string
	arr[0] = "posição 1"
	fmt.Println(arr)

	arr2 := [5]string{"posiçao1", "posiçao2", "posiçao3", "posiçao4", "posiçao5"}
	fmt.Println((arr2))

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println((arr3))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

}
