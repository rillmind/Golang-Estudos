package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escreva um programa que conte o número de letras de uma frase
// informada pelo usuário.

func main() {
	var (
		letras int
		frase  string
		s      = bufio.NewScanner(os.Stdout)
	)

	fmt.Print("Escreva uma frase: ")
	s.Scan()
	frase = s.Text()

	for _, byte := range frase {
		// Esses intervalos são os intervalos de letras maiúsculas e minúsculas da tabelas ascii respectivamentes
		if (byte >= 65 && byte <= 90) || (byte >= 97 && byte <= 122) {
			letras++
		}
	}

	fmt.Println(letras)
}
