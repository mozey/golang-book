package main

import "fmt"
import "errors"

// Go has a built-in type for errors (the error type),
// we can create our own errors by using the New function in the errors package
func main() {
	err := errors.New("error message")
	fmt.Println(err)
}