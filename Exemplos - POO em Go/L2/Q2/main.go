package main

import "fmt"

// Escreva um programa que leia 10 números e mostre as seguintes informações:
// a) Menor valor digitado
// b) Maior valor digitado
// c) Média dos valores digitados

func main() {
	var (
		menor, maior, soma, numero int
	)

	for range 10 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		switch {
		case menor == 0 || maior == 0:
			menor = numero
			maior = numero
		case numero < menor:
			menor = numero
		case numero > maior:
			maior = numero
		}

		soma += numero
	}

	fmt.Println("O menor número é:", menor)
	fmt.Println("O maior número é:", maior)
	fmt.Println("A média dos números é:", soma/10)
}
