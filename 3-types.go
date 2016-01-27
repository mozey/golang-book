package main

// Import multiple packages like this
import (
	"fmt"
)

func main() {
	// Numbers .................................................................
	// 1 + 1 is an expression:
	// (numeric literal) operator (numeric literal)

	// Go has the following operators
	// 		+	addition
	// 		-	subtraction
	// 		*	multiplication
	// 		/	division
	// 		%	remainder (mod)

	fmt.Println("int: 1 + 1 =", 1 + 1)

	// Use the .0 to tell Go that this is a floating point number
	fmt.Println("float: 1 + 1 =", 1.0 + 1.0)
	fmt.Println()


	// Strings .................................................................

	// Two ways to create a string literal
	one := "Double quote string with tab \t and new \nline escape sequences."
	two := `Back tick strings
can span multiple lines`
	fmt.Println(one, two)
	fmt.Println()

	// Find the length of a string with len()
	s := "Hello World"
	fmt.Println(s, "is", len(s), "character long")
	fmt.Println()

	// Strings are "indexed" starting at 0 not 1.
	// Character is represented as a byte, and byte is an integer.
	char := "Hello World"[1]
	fmt.Println(
		"According to the ASCII table;",
		fmt.Sprintf("%d in decimal =", char),
		fmt.Sprintf("%x in hex =", char),
		"the letter", string(char))
	fmt.Println()

	// Concatenation uses the same symbol as addition.
	// The Go compiler figures out what to do based on the types of the arguments.
	fmt.Println("This is a " + "concatenated string")
	fmt.Println()

	// Booleans ................................................................

	// Three logical operators are used with boolean values
	// 		&&	and
	// 		||	or
	// 		!	not

	// Note that the bitwise XOR operator (int ^ int) applies only to integers.

	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

}

