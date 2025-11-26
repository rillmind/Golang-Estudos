package main

import "fmt"

// Escreva um programa que leia a idade e o nome de 10 pessoas. Escreva métodos para:
// • Retornar o nome da pessoa mais velha;
// • Retornar o nome da pessoa mais nova;
// • O numero de pessoas com mais de 18 anos

func main() {
	var (
		pessoas []pessoa
		pessoa  pessoa
	)

	for i := range 3 {
		fmt.Printf("Nome da %d° pessoa: ", i+1)
		fmt.Scan(&pessoa.nome)

		fmt.Printf("Idade da %d° pessoa: ", i+1)
		fmt.Scan(&pessoa.idade)

		pessoas = append(pessoas, pessoa)
	}

	pessoaMaisVelha := pessoaMaisVelha(pessoas)
	fmt.Printf("A pessoa mais velha é %s\n", pessoaMaisVelha.nome)

	pessoaMaisNova := pessoaMaisNova(pessoas)
	fmt.Printf("A pessoa mais nova é %s\n", pessoaMaisNova.nome)

	maiores := pessoasDeMaior(pessoas)
	fmt.Printf("Total de pessoas maiores de idade: %d\n", maiores)
}
