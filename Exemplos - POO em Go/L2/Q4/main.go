package main

import "fmt"

// Escreva um programa que leia 10 números e mostre na tela números na
// ordem inversa.

func main() {
	var (
		numero  int
		numeros []int
	)

	for range 10 {
		fmt.Print("Digite um numero: ")
		fmt.Scan(&numero)

		numeros = append(numeros, numero)
	}

	for i := 9; i >= 0; i-- {
		fmt.Println(numeros[i])
	}
}
