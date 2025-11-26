package main

import "fmt"

// Escreva um programa que leia dois vetores de 05 elementos cada. Após ler
// todos os elementos, verifique se os vetores são iguais e mostre na tela.

func main() {
	var (
		ns1, ns2 []int
		number   int
		equals   bool = true
	)

	for range 5 {
		fmt.Print("Primeiro slice: ")
		fmt.Scan(&number)

		ns1 = append(ns1, number)
	}

	for range 5 {
		fmt.Print("Segundo slice: ")
		fmt.Scan(&number)

		ns2 = append(ns2, number)
	}

	for i := range 5 {
		if ns1[i] != ns2[i] {
			fmt.Println("Slices diferentes")
			equals = false
			return
		}
	}

	if equals {
		fmt.Println("Slices identicos")
	}
}
