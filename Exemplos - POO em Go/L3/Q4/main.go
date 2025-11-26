package main

import (
	"fmt"
	"strings"
)

// Escreva um programa com um método que receba o salário de um
// funcionário e calcule o imposto que ele deve pagar. Utilize os seguintes valores para
// calcular o imposto:
// • Até R$ 2.000,00 – Isento de imposto
// • De R$ 2.000,01 até R$ 3.500,00 – 15% de imposto
// • De R$ 3.500,01 até R$ 5.000,00 – 22% de imposto
// • Acima de R$ 5.000,01 – 30% de imposto

func main() {
	imposto(2000)
	imposto(3000)
	imposto(3550)
	imposto(5500)
}

func imposto(salario float32) {
	switch {
	case salario <= 2000:
		salarioStr := fmt.Sprintf("Salário: R$%.2f. Isento de imposto.\n", salario)
		salarioStr = strings.Replace(salarioStr, ".", ",", 1)
		fmt.Println(salarioStr)
	case salario > 2000 && salario <= 3500:
		calcular(salario, 15)
	case salario > 3500 && salario <= 5000:
		calcular(salario, 22)
	case salario > 5000:
		calcular(salario, 30)
	}
}

func calcular(salario, imposto float32) {
	salarioStr := fmt.Sprintf("Salário: R$%.2f. Imposto de %.0f porcento.\n", salario-((imposto/100.00)*salario), imposto)
	salarioStr = strings.Replace(salarioStr, ".", ",", 1)
	fmt.Println(salarioStr)
}
