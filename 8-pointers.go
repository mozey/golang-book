package main

import "fmt"

// Function parameters are passed by value,
// that is, the argument x is copied to the function.
func zeroByValue(x int) {
	x = 0
}

// Use pointers to modify the original x variable in the main function.
func zeroByReference(xPtr *int) { // xPtr is a pointer to an int
	// Store the int 0 in the memory location xPtr refers to
	// Assigning 0 to xPtr instead of *xPtr results in an error:
	// "cannot use 0 (type int) as type *int in assignment"
	fmt.Println("xPtr =", xPtr)
	fmt.Println("*xPtr =", *xPtr)
	*xPtr = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *float64) {
	*x = *x * *x
}

func main() {
	x := 5
	zeroByValue(x)
	fmt.Println("zeroByValue =>", x) // x is still 5

	zeroByReference(&x)
	fmt.Println("zeroByReference =>", x) // x is 0

	xFloat := 1.5
	square(&xFloat)
	fmt.Println("square =>", xFloat) // x is 2.25
}

