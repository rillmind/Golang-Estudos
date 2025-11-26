package main

import "fmt"

// Escreva um programa que leia a altura de 05 pessoas e calcule a média.

func main() {
	var soma, peso float32

	for range 5 {
		fmt.Print("Digite um peso: ")
		fmt.Scan(&peso)
		soma += peso
	}

	fmt.Printf("%0.2f\n", soma/5)
}
