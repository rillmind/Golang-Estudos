package main

import "fmt"

// Escreva um programa que leia 10 números e armazene em um vetor. Em
// seguida, mostre a soma desses números.

func main() {
	var (
		numbers      []int
		number, soma int
	)

	for range 10 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}

	for _, number := range numbers {
		soma += number
	}

	fmt.Println("Soma:", soma)
}
