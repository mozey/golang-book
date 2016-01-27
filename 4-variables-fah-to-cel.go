package main

import (
	"fmt"
)

func main() {
	// Convert fahrenheit to celsius
	fmt.Print("Enter a fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5/9
	fmt.Println(celsius)
}