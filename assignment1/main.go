package main

import "fmt"

func main() {
	intSlice := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, myInt := range intSlice {
		if myInt%2 == 0 {
			fmt.Printf("%v is even\n", myInt)
		} else {
			fmt.Printf("%v is odd\n", myInt)
		}

	}
}
