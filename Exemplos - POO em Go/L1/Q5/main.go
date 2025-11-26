package main

import "fmt"

// Escreva um programa que leia dois números e a operação aritmética desejada.

// Se o usuário digita ‘+’, retorne o resultado da soma entre os números;
// Se o usuário digita ‘-’, retorne o resultado da subtração entre os números;
// Se o usuário digita ‘*’, retorne o resultado da multiplicação entre os números;
// Se o usuário digita ‘/’, retorne o resultado da divisão entre os números;

func main() {
	var (
		n1, n2    int
		operation string
	)

	fmt.Print("Número 1: ")
	fmt.Scan(&n1)

	fmt.Print("Número 2: ")
	fmt.Scan(&n2)

	fmt.Print("Operação: ")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println(n1 + n2)

	case "-":
		fmt.Println(n1 - n2)

	case "*":
		fmt.Println(n1 * n2)

	case "/":
		fmt.Println(n1 / n2)

	default:
		fmt.Println("Operação inválida.")
	}
}
