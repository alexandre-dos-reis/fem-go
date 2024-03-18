package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}

	// Add data to channel
	channel <- "Done !"
}

func main() { // main goroutine

	var channel chan string

	go printMessage("Frontend masters rocks !", channel)
	// Wait for the value before continuing execution
	response := <-channel

	fmt.Println(response)
	// !! Don't add goroutine to the last func because main goroutine will end first rendering nothing !!
}
