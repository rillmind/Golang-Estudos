package main

import "fmt"

// Write a function called prefixer that has an input parameter of type string
// and returns a function that has an input parameter of type string and returns
// a string. The returned function should prefix its input with the string passed
// into prefixer. Use the following main function to test prefixer:

func main() {
	x := 10
	pointerX := &x
	fmt.Println(pointerX)
	fmt.Println(*pointerX)

	*pointerX = 5 + *pointerX
	fmt.Println(x)
}
