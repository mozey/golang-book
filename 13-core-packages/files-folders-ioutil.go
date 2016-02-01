package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("13-core-packages/io.go")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}