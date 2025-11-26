package main

import "fmt"

// Escreva um programa com um método que receba um vetor como
// parâmetro e calcule a média dos valores contidos no vetor. Obs.: O tamanho do vetor
// deve ser informado pelo usuário.

func main() {
	var (
		s []float32
		n float32
	)

	fmt.Println("0 para sair.")

	for {
		fmt.Print("Digite um número real: ")
		fmt.Scan(&n)
		fmt.Println()

		s = append(s, n)

		if n == 0 {
			break
		}
	}

	calcularMediaDoSlice(s)
}

func calcularMediaDoSlice(s []float32) {
	var (
		soma float32
	)

	for _, n := range s {
		soma += n
	}

	fmt.Printf("A média dos números é: %.2f\n", soma/float32(len(s)))
}
