package main

import "fmt"

// math is the name of a package that is part of Go's standard distribution.
import "math"

// Ours package name is "github.com/mozey/golang-book/11-packages/math"
//import "github.com/mozey/golang-book/11-packages/math"
// Using an alias "m" lets us use both math packages.
import m "github.com/mozey/golang-book/11-packages/math"

func main() {
	pi := math.Pi
	fmt.Println(pi)

	xs := []float64{1, 2, 3, 4}
	//avg := math.Average(xs)
	avg := m.Average(xs)
	fmt.Println(avg)

	min := m.Min(xs)
	max := m.Max(xs)
	fmt.Println(min)
	fmt.Println(max)
}