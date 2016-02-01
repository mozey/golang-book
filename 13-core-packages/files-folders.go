package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("13-core-packages/io.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Make sure the file is closed as soon as the function completes.
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}