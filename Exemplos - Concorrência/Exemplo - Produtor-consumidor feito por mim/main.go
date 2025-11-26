package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rillmind/colors"
)

const (
	numberOfPizzas = 10
)

var (
	pizzasMade, pizzasFail, total int
)

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber uint8
	message     string
	sucess      bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber uint8) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		sucess := false

		if rnd < 5 {
			pizzasFail++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making pizza #%d! It will take %d seconds...\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		switch {
		case rnd <= 2:
			msg = fmt.Sprintf("*** We ran out of engridients for pizza #%d!", pizzaNumber)

		case rnd <= 4:
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)

		default:
			sucess = true
			msg = fmt.Sprintf("Pizza order #%d! is ready!", pizzaNumber)
		}

		pizza := PizzaOrder{
			pizzaNumber,
			msg,
			sucess,
		}

		return &pizza
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of what of which pizza we are makin
	var (
		i uint8
	)
	//run forever until we receive a quit notification
	// try to make the pizzas
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	colors.Green("Heyy")
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out message
	colors.Sapphire("The Pizzeria its open for business!")

	// create a producer
	pizzaJob := &Producer{
		make(chan PizzaOrder),
		make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run costumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.sucess {
				colors.Green(i.message)
				colors.Green(fmt.Sprintf("Order #%d is out for delivery!", i.pizzaNumber))
				fmt.Println()
			} else {
				colors.Red(i.message)
				colors.Red("The costumer is really mad!")
			}
		} else {
			colors.Sapphire(("Done of making pizzas..."))

			err := pizzaJob.Close()
			if err != nil {
				colors.Red(fmt.Sprintf("*** Error closing channel! %s", err))
			}
		}
	}

	// print out ending message
}
