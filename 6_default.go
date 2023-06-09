package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	fmt.Println("the program started")

	for {
		go func() {
			time.Sleep(1 * time.Second)
			ch1 <- "Hello"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			ch2 <- "World"
		}()

		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		default:
			fmt.Println("No messages received")
		}
	}
}
