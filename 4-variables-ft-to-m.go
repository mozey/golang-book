package main

import (
	"fmt"
)

func main() {
	// Convert feet to meters
	fmt.Print("Enter feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := 0.3048 * feet
	fmt.Println(meters)
}