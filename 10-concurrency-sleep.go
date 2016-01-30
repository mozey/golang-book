package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	select {
		case <-time.After(time.Duration(seconds) * time.Second):
			return
	}
}

func main() {
	Sleep(3)
	fmt.Println("Done sleeping")
}