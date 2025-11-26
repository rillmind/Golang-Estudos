package main

import (
	"fmt"
	"strings"
)

// Escreva um programa com um método que receba uma String e retorne o
// número de vogais dessa String.

func main() {
	vogais("raios eh muito poderoso")
}

func vogais(text string) {
	var (
		letra  string
		soma   int
		vogais = []string{"A", "E", "I", "O", "U"}
	)

	for _, byte := range text {
		letra = strings.ToUpper(string(byte))

		for _, vogal := range vogais {
			if letra == vogal {
				soma++
			}
		}
	}

	fmt.Println(soma)
}
