package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que leia uma palavra e mostre a quantidade de vogais da palavra.

func main() {
	var (
		frase, letra string
		soma         int
		sc           = bufio.NewScanner(os.Stdout)
		vogais       = []string{"A", "E", "I", "O", "U"}
	)

	fmt.Print("Digite uma frase: ")
	sc.Scan()
	frase = sc.Text()

	for _, byte := range frase {
		letra = strings.ToUpper(string(byte))

		for _, vogal := range vogais {
			if letra == vogal {
				soma++
			}
		}
	}

	fmt.Println(soma)
}
