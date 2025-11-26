package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 1 {
		return n
	} else {
		return n + fibonacci(n-n)
	}
}

func main() {
	fmt.Println(fibonacci(4))
}
