package main

import "fmt"

// Take in the argument named xs, a slice of float64s, and return one float64
//func average(xs []float64) float64 {
func average(xs []float64) (r float64) {
	// panic causes a run time error
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	// Return causes the function to immediately stop
	// and return the value to the calling function.
	//return total / float64(len(xs))

	// The return value can be named
	r = total / float64(len(xs))
	return
}

// Functions can return multiple values,
// this is often used to return an error value...
func f1() (string, error) {
	return "error example", nil
}

// ...or success flag
func f2() (string, bool) {
	return "success flag example", true
}

// Variadic function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Closures can be used to create generator functions
func makeNumberGenerator(init uint) func() uint {
	i := init
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}


func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func half(x int) (int, bool) {
	even := x % 2 == 0
	return x/2, even
}

func main() {
	// The parameter name does not have to match the function signature
	someOtherName := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(someOtherName))

	// Function can return multiple values
	s, err := f1()
	fmt.Println(s, err)

	// We can use the new variable syntax
	// if there is at least one new variable on the LHS.
	s, ok := f2()
	fmt.Println(s, ok)

	//fmt.Println(add(1,2,3))
	xs := []int{1, 2, 3}
	// We also can pass in slice
	fmt.Println(add(xs...))


	// Closure .................................................................
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeNumberGenerator(uint(0))
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	nextOdd := makeNumberGenerator(uint(1))
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3


	// Recursion ...............................................................

	fmt.Println(factorial(5))


	// Defer ...................................................................

	defer second()
	first()


	// Panic and Recover .......................................................

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

	// .........................................................................

	fmt.Println(half(2))
	fmt.Println(half(1))
}

