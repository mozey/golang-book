package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name type is assigned the string "Hello World"
	var x string = "Hello World"
	var y string
	y = "!"

	// The assignment statement is shorthand for x = x + y
	x += y
	fmt.Println(x)

	// The equality operator returns a boolean
	fmt.Println(x == y)

	// The compiler will infer type if none is specified
	var a = "string"
	// := is shorthand for creating a new variable with a starting value.
	b := 123
	fmt.Println("a is type", reflect.TypeOf(a))
	fmt.Println("b is type", reflect.TypeOf(b))

	const c = "cannot be chaged"
	 //c = "Abc" // This result in a compile-time error

	// Use the keyword var (or const) followed by parentheses
	// with each variable on its own line to declare multiple variables.
	var (
		multiPass = "access"
		mr = "Zorg"
		perfect = "being"
	)
	fmt.Println(multiPass, mr, perfect)

	// Take in a number entered by the user and double it
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}