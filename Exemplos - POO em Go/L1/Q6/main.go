package main

import "fmt"

// Leia o peso e a altura de 05 pessoas e calcule o IMC e armazene em um
// vetor. Após o termino do programa, imprima as seguintes informações:

// O Número de pessoas que estão abaixo do peso (IMC < 18,5);
// O Número de pessoas que estão dentro do peso ideal (IMC > 18,5 e IMC < 25);
// O Número de pessoas que estão acima do peso (IMC > = 25)

func main() {
	var (
		imcs                 []float32
		peso, altura, imc    float32
		abaixo, acima, media int
	)

	for range 5 {
		fmt.Print("Digite um peso: ")
		fmt.Scan(&peso)

		fmt.Print("Digite uma altura: ")
		fmt.Scan(&altura)

		imc = peso / (altura * altura)

		imcs = append(imcs, imc)
	}

	for _, imc := range imcs {
		switch {
		case imc < 18.5:
			abaixo++
		case imc > 18.5 && imc < 25:
			media++
		case imc >= 25:
			acima++
		}
	}

	fmt.Println(abaixo, "pessoas estão abaixo do peso ideal.")
	fmt.Println(media, "pessoas estão no peso ideal.")
	fmt.Println(acima, "pessoas setão acima do peso ideal.")
}
