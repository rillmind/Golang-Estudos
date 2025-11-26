package main

type pessoa struct {
	nome  string
	idade int
}

func pessoaMaisVelha(pessoas []pessoa) *pessoa {
	var (
		pessoaMaisVelha *pessoa
	)

	for i, p := range pessoas {
		switch {
		case i == 0:
			pessoaMaisVelha = &p

		case pessoas[i].idade > pessoaMaisVelha.idade:
			pessoaMaisVelha = &p
		}
	}

	return pessoaMaisVelha
}

func pessoaMaisNova(pessoas []pessoa) *pessoa {
	var (
		pessoaMaisNova *pessoa
	)

	for i, p := range pessoas {
		switch {
		case i == 0:
			pessoaMaisNova = &p

		case pessoas[i].idade < pessoaMaisNova.idade:
			pessoaMaisNova = &p
		}
	}

	return pessoaMaisNova
}

func pessoasDeMaior(pessoas []pessoa) int {
	var (
		soma int
	)

	for i := range pessoas {
		switch {
		case pessoas[i].idade >= 18:
			soma++
		}
	}

	return soma
}
