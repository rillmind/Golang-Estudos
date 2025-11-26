package main

import (
	"Q1/student"
	"fmt"
	"time"
)

// Crie uma classe estudante. Esta classe deve ter nome, número de
// matrícula, endereço, e 04 notas. Forneça um construtor e métodos para alterar e
// obter o nome do estudante e as notas. Além disso, crie um método para calcular a
// média do estudante. Crie outra classe que forneça as seguintes opções ao usuário:
// 1. Criar um Estudante
// 2. Calcular Média (Se o Aluno foi Aprovado, informe ao usuário)
// 3. Obter número de matricula
// 4. Obter Endereço
// 0. Sair

func main() {
	var (
		s      student.Student
		option uint8
	)

loop:
	for {
		menu()
		fmt.Print("\nEscolha: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Nome do aluno: ")
			fmt.Scan(&s.Name)
			fmt.Print("Número de matrícula: ")
			fmt.Scan(&s.RNumber)
			fmt.Print("CEP: ")
			fmt.Scan(&s.Adress.CEP)
			fmt.Print("Número da casa: ")
			fmt.Scan(&s.Adress.Number)
			for i := range s.PNotes {
				fmt.Printf("Nota %d: ", i+1)
				fmt.Scan(&s.PNotes[i])
			}

		case 2:
			average := s.Average()
			if average >= 7 {
				fmt.Println("Aprovado!")
			}

			fmt.Println("Média do aluno:", average)

		case 3:
			fmt.Println("Número de matrícula:", s.RNumber)

		case 4:
			fmt.Println(s.Adress.ToJSON())

		case 0:
			fmt.Println("Saindo...")
			time.Sleep(time.Millisecond * 500)
			break loop

		default:
			fmt.Println("Opção inválida")
		}
	}
}

func menu() {
	print(`
1. Criar um Estudante
2. Calcular Média (Se o Aluno foi Aprovado, informe ao usuário)
3. Obter número de matricula
4. Obter Endereço
0. Sair
`)
}
