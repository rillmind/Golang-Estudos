package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string) {
	for {
		s := <-ping

		ping <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)

	go shout(ping)

	fmt.Println("Digite uma frase ('q' para sair)")

	for {
		fmt.Print("> ")

		userInput := ""
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput

		res := <-ping

		fmt.Println(res)
	}

	fmt.Println("Fechando canal")

	close(ping)
}
