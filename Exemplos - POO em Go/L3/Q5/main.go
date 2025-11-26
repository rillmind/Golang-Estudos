package main

import "fmt"

// Escreva um programa com métodos para calcular a área (b * h) e o
// perímetro de um retângulo (2b + 2h). Forneça as seguintes opções ao usuário:
// 1. Calcular a área do retângulo
// 2. Calcular o perímetro do retângulo

func main() {
	area(3, 2)
	perimetro(3, 2)
}

func area(b, h float32) {
	fmt.Println(b * h)
}

func perimetro(b, h float32) {
	fmt.Println((b * 2) + (h * 2))
}
