package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("ola mundo", channel)
	fmt.Println("depois da func write")

	for message := range channel {
		fmt.Println(message)
	}
	// for {
	// 	msg, Open := <-channel
	// 	if !Open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
}

func write(text string, channel chan string) {

	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}

//caso de deadlock é pq o canal ta esperando um valor que nunca vai chegar, e esse erro nao é pego na hora da compilação
