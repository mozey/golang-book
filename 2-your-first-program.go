// This is a comment

// Executables always start with "package main"
// Libraries start with "package libname"
package main // The package declaration

// Importing the library "fmt"
// Text wrapped in double quotes are "string literals",
// they are a type of "expression"
import "fmt"

// "main" is a special function name,
// it is called when you execute the program
func main() {
	// Call the function Println in package fmt with argument "Hello World"
	// To get documentation for fmt.Println run this command:
	// godoc fmt Println
	fmt.Println("Hello World")
	fmt.Println("Hello, my name is " + "Chris")
}

