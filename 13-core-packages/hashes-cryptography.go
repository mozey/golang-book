package main

// Hash functions in Go are broken into two categories:
// cryptographic and non-cryptographic.

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main() {
	h := crc32.NewIEEE()
	// The crc32 hash object implements the Writer interface,
	// so we can write bytes to it like any other Writer
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("13-core-packages/error.go")
	if err != nil {
		return
	}
	h2, err := getHash("13-core-packages/io.go")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}