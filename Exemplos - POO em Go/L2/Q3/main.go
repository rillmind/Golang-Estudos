package main

import "fmt"

// Leia as notas de 10 alunos e armazene em um vetor. Calcule a media das
// notas informadas e mostre quantos alunos ficaram acima da média.

func main() {
	var (
		i          int
		notas      []float32
		nota, soma float32
	)

	for range 10 {
		fmt.Print("Digite a nota do aluno: ")
		fmt.Scan(&nota)

		notas = append(notas, nota)
	}

	for _, nota := range notas {
		if nota >= 7 {
			i++
		}

		soma += nota
	}

	fmt.Println("A média das notas foi:", soma/10)
	fmt.Println(i, "alunos ficaram acima da média")
}
