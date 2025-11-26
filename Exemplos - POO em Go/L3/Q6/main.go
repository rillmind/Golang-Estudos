package main

import (
	"fmt"
	"log"
	"time"
)

// Escreva um programa com métodos que façam a conversão de Fahrenheit
// para Celsius (C = 5(F - 32) / 9) e de Kelvin para Celsius (C = (K - 273)). Forneça as
// seguintes opções ao usuário:
// 1. Transformar de Fahrenheit para Celsius
// 2. Transformar de Kelvin para Celsius

func main() {
	var (
		f, k    float32
		escolha int
	)

	for {
		menu()

		fmt.Print("O que deseja converter? ")
		fmt.Scan(&escolha)
		fmt.Println()

		switch escolha {
		case 1:
			fmt.Print("Temperatura em Fahrenheit: ")
			fmt.Scan(&f)
			fahrenheit(f)

		case 2:
			fmt.Print("Temperatura em Kelvin: ")
			fmt.Scan(&k)
			kelvin(k)

		case 0:
			log.Fatal("Saindo...")
			time.Sleep(time.Millisecond * 500)

		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func fahrenheit(f float32) {
	fmt.Println(5 * (f - 32) / 9)
}

func kelvin(k float32) {
	fmt.Println(k - 273)
}

func menu() {
	fmt.Printf(`
1. Converter de Fahrenheit para Celsius
2. Converter de Kelvin para Celsius
0. Sair

`)
}
