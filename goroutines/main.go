package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("ola mundo")
		waitGroup.Done()
	}()

	go func() {
		write("ola mundo2")
		waitGroup.Done()
	}()
	waitGroup.Wait()

}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
