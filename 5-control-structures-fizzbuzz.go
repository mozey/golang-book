package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        s := fmt.Sprintf("%v", i)
        if i % 3 == 0 {
			if (i % 5 == 0) {
				s = "FizzBuzz"
			} else {
				s = "Fizz"
			}
        } else if i % 5 == 0 {
			if (i % 3 == 0) {
				s = "FizzBuzz"
			} else {
				s = "Buzz"
			}
        }

		if (s == fmt.Sprintf("%v", i)) {
			fmt.Println(s, "\t-")
		} else {
			// Printing the number in front is not the spec but whatever...
			if (i == 100) {
				fmt.Println(i, "-", s)
			} else {
				fmt.Println(i, "\t-", s)
			}
		}
    }
}


