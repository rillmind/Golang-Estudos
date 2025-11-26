package main

import "fmt"

// Escreva um programa com um método que receba um número e determine
// se esse número é par ou ímpar.

func main() {
	verify(2)
	verify(3)
}

func verify(n int) {
	if n%2 == 0 {
		fmt.Println("Par")
		return
	}
	fmt.Println("Impar")
}
