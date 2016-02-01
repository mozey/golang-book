package main

import "fmt"

func main() {
    // The io package consists mostly of interfaces used in other packages.
	// The two main interfaces are Reader and Writer,
	// the first has a Read and the latter a Write method.

	// io.Copy
	// func Copy(dst Writer, src Reader) (written int64, err error)

	// bytes.Buffer
	// Used to read or write to a []byte or a string

	// buf.Bytes() to convert a buffer to []byte

	// When reading from strings only, use strings.NewReader for efficiency.

	fmt.Println("TODO Make some examples")
}