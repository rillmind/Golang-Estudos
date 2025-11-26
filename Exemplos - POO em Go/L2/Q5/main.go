package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que leia o nome de X pessoas. Após ler todos os
// nomes, pergunte ao usuário o nome que ele deseja consultar e verifique se este nome
// consta na lista.
// Obs.: O número de pessoas deve ser determinado pelo usuário.

func main() {
	var (
		nome, busca, quer string
		nomes             []string
		qnt               int
	)

	fmt.Print("Quantas pessoas deseja adicionar a lista? ")
	fmt.Scan(&qnt)

	for range qnt {
		fmt.Print("Adicione um nome a lista:")
		fmt.Scan(&nome)

		nomes = append(nomes, nome)
	}

	fmt.Print("Deseja buscar alguem na lista? [S/N] ")
	fmt.Scan(&quer)

	querM := strings.ToUpper(quer)

	if strings.Contains("SIM", querM) {
		fmt.Print("Nome que deseja buscar: ")
		fmt.Scan(&busca)

		for _, nome := range nomes {
			if nome == busca {
				fmt.Println("Usuario encontrado!")
				return
			} else {
				fmt.Println("Usuario nao encontrado!")
				return
			}
		}
	}
}
