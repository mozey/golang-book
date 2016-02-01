package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})

	// There is no native type to represent a 160 bit number,
	// so we use a slice of 20 bytes instead
	fmt.Println(bs)

	// To print the string equivalent of the []byte
	fmt.Printf("%x\n", bs)
}