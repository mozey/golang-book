package main

// In addition to lists and maps Go has several more
// collections available underneath the container package.
// The container/list package implements a doubly-linked list
import (
	"fmt"; "container/list"
)


func main() {
	var x list.List

	// The zero value for a List is an empty list
	// (a *List can also be created using list.New)
	fmt.Println(x)

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
