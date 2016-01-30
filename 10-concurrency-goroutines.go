package main

import (
	"fmt"
	"time"
	"math/rand"
)


func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		// Sleep somewhere between 0 and 250ms
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// This program consists of two goroutines.
// The the main function itself is an implicit goroutine.
func main() {
	// The second goroutine is created with the go keyword
	//go f(10)
	for i := 0; i < 10; i++ {
    	go f(i)
  	}

	// Calling a goroutine returns immediately to the next line without waiting,
	// fmt.Scanln is included so the numbers can print before the program exits.
	var input string
	fmt.Scanln(&input)
}

