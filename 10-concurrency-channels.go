package main

import (
	"fmt"
	"time"
)

func pinger(c <-chan string) {
	for i := 0;; i++ { // Loop forever
		// Send a message on the channel
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0;; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for { // Another way to loop forever
		// Receive a message from the channel
		//msg := <-c

		// Print the channel message and sleep for one second
		//fmt.Println(msg)
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Use the chan keyword followed by a type to create a channel
	var c chan string = make(chan string)

	// Run the senders and receiver in goroutines.
	// Using a channel like this synchronizes the goroutines.
	// The senders will wait until the receiver is ready to print.
	go pinger(c)
	go ponger(c)
	go printer(c)

	// Hit enter to stop the program
	var input string
	fmt.Scanln(&input)
}