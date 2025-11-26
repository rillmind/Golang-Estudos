package main

import (
	"fmt"
)

func main() {
	intSlice := []int{7, 4, 2, 5, 6, 7, 8, 1, 12, 42, 52, 5}

	for _, number := range intSlice {
		if number%2 == 0 {
			fmt.Printf("%v is even\n", number)
		} else {
			fmt.Printf("%v is odd\n", number)
		}
	}
}
