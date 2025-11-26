package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// Escreva um programa que leia 02 números a operação aritmética desejada. Se o usuário digitar:
//
// 1. Retornar a soma entre dois números;
// 2. Retornar a subtração entre dois números;
// 3. Retornar a multiplicação entre dois números;
// 4. Retornar a divisão entre dois números;
//
// Obs.: Cada opção deve ser um método diferente.

func main() {
	var (
		n1, n2, operacao int
	)

	fmt.Print("Primeiro número: ")
	fmt.Scanln(&n1)

	fmt.Print("Segundo número: ")
	fmt.Scan(&n2)

	for {
		cmd := exec.Command("clear")

		cmd.Run()

		menu()

		fmt.Print("Operação: ")
		fmt.Scan(&operacao)

		switch operacao {
		case 1:
			somar(n1, n2)

		case 2:
			subtrair(n1, n2)

		case 3:
			multiplicar(n1, n2)

		case 4:
			dividir(n1, n2)

		case 0:
			time.Sleep(time.Second)
			log.Fatalln("Saindo...")

		default:
			log.Fatalln("Operação inválida!")
		}
	}
}

func somar(n1, n2 int) {
	fmt.Println(n1 + n2)
}

func subtrair(n1, n2 int) {
	fmt.Println(n1 - n2)
}

func multiplicar(n1, n2 int) {
	fmt.Println(n1 * n2)
}

func dividir(n1, n2 int) {
	fmt.Println(n1 / n2)
}

func menu() {
	fmt.Print(`
1. Soma
2. Subtração
3. Multiplicação
4. Divisão
0. Sair

`)
}
