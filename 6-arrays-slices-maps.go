package main

import "fmt"

func main() {
	// Arrays ..................................................................
	// Array are indexed starting at zero
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// Suppose we want to compute an average grade...
	// This is the verbose method of initialising an array with values
	//var y [5]float64
	//y[0] = 98
	//y[1] = 93
	//y[2] = 77
	//y[3] = 82
	//y[4] = 83
	// The shorthand is much more concise, and it can span multiple lines
	y := [5]float64{
		98, 93, 77, 82, 83, // The trailing comma is required
	}

	// This is bad code because the array length is hard-coded
	//var total float64 = 0
	//for i := 0; i < 5; i++ {
	//	total += y[i]
	//}
	//fmt.Println(total / 5)

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}

	// Will error with "mismatched types float64 and int".
	// The len function returns an int
	//fmt.Println(total / len(y))

	// We can use "type conversion" to convert the int to a float64
	fmt.Println(total / float64(len(x)))

	// The range keyword lets us use a special form of for loop.
	// Will error with "i declared and not used"
	//for i, value := range x {
	// Use underscore to discard unused return values
	total = 0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))


	// Slices ..................................................................
	// Create a slice with length 0
	var s []float64
	fmt.Println("s", s)

	// Create a slice that is associated with a float64 array of length 5.
	//m := make([]float64, 5)
	// The third parameter represents the capacity of the underlying array.
	n := make([]float64, 5, 10)
	fmt.Println("n", n, "len(n) =", len(n))

	// Slices can also be created using a range expression
	// arr[0:]	=> arr[0:len(arr)]
	// arr[:5]	=> arr[0:5]
	// arr[:]	=> [0:len(arr)]
	arr := [5]float64{1, 2, 3, 4, 5}
	slice := arr[0:3]
	fmt.Println("slice", slice)


	// Slice functions .........................................................

	slice1 := []int{1, 2, 3}
	// Create new slice from slice1 and the following arguments
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	// Copy as many elements as possible from slice3 to slice4
	copy(slice4, slice3)
	fmt.Println(slice4)
	fmt.Println(slice3, slice4)


	// Maps ....................................................................
	// Create a map m,
	// m is a map of strings to ints

	// This code causes a runtime error "assignment to entry in nil map",
	// maps have to be initialized before they can be used.
	//var m map[string]int
	//m["key"] = 10

	m := make(map[string]int)
	m["key"] = 10
	m["foo"] = 11
	delete(m, "foo") // Delete a key from the map
	fmt.Println(m)

	// For undefined keys maps returns the zero value for the value type.
	if foo, keyFound := m["foo"]; keyFound {
		fmt.Println("foo is", foo)
	} else {
		fmt.Println("Key \"foo\" not found")
	}

	// Shorthand for creating maps
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
	}
	fmt.Println(elements)
}

