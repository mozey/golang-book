package main

import "fmt"

func main() {
	// The commented code is equivalent to the loop below
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i = i + 1
	//}
	for i := 1; i <= 4; i++ { // Special operator ++ increments a variable
		s := fmt.Sprintf("%v", i)

		if i % 2 == 0 {
			s += " even"
		} else {
			s += " odd"
		}

		if i % 2 == 0 {
			s += " divisible by 2"
		} else if i % 3 == 0 {
			s += " divisible by 3"
		} else if i % 4 == 0 {
			s += " divisible by 4"
		}

		fmt.Println(s)
	}

	arr := [3]int{1, 2}
	for _, value := range arr {
		fmt.Println(value)
	}

	j := 3
	switch j {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}
}