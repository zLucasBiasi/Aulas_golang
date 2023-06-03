package main

import "fmt"

func recoverExection() {
	if r := recover(); r != nil {
		fmt.Println("recuperamo familia")
	}
}

func isAproved(n1, n2 float64) bool {
	defer recoverExection()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("a media é exatamente 6")
}

func main() {
	fmt.Println(isAproved(6, 6))
	fmt.Println("executou")
}
