// go run cli-arguments.go -max=100
package main

import (
	"fmt"; "flag"; "math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))

	// Any additional non-flag arguments can be retrieved
	// with flag.Args() which returns a []string.
}
