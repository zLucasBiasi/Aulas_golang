package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "ola mundo"
	channel <- "ola mundo2"
	// channel <- "ola mundo" (aqui chegou na capacidade maxima ai da o erro de deadlock)
	message := <-channel
	fmt.Println(message)
}
