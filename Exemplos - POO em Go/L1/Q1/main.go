package main

import "fmt"

// Escreva um programa que calcule a idade do usuário. Para isso, o
// programa deve ler o ano de nascimento, o ano atual e depois mostrar a idade.
// Obs.: O ano deve ser informado no formato AAAA (ex.: 1980).

func main() {
	var birth, currYear int

	fmt.Print("Ano de nascimento: ")
	fmt.Scan(&birth)

	fmt.Print("Ano atual: ")
	fmt.Scan(&currYear)

	fmt.Println(currYear - birth)
}
