package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escreva um programa que inverta uma frase informada pelo usuário e
// mostre na tela.

func main() {
	var (
		frase string
		s     = bufio.NewScanner(os.Stdout)
	)

	fmt.Print("Digite uma frase: ")
	s.Scan()
	frase = s.Text()

	for i := len(frase) - 1; i >= 0; i-- {
		fmt.Print(string(frase[i]))
	}
}
