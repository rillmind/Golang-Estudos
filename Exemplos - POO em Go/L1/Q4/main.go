package main

import "fmt"

// Escreva um programa que leia a idade de várias pessoas e calcule a média.
// O programa deve ler as idades até o usuário digitar 0 (zero).

func main() {
	var i, idade, soma int

	fmt.Println("Digite 0 para sair.")

	for {
		fmt.Println("Digite a ideia de uma pessoa: ")
		fmt.Scan(&idade)

		if idade == 0 {
			break
		}

		soma += idade
		i++
	}

	fmt.Printf("Média das idades: %d\n", soma/i)
}
