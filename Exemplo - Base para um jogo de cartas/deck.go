package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	deck := deck{}
	cardSuits := [4]string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades",
	}
	cardValues := [13]string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, card := range cardValues {
			deck = append(deck, card+" of "+suit)
		}
	}

	return deck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%02v: %v\n", i, card)
	}
}

func (d deck) toString() string {
	var deck string
	for i, card := range d {
		if i == len(d)-1 {
			deck += card
		} else {
			deck += card + ","
		}
	}
	return deck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")

	return deck(stringSlice)
}

func (d deck) shuffle() {
	currentTime := time.Now().UnixNano()
	source := rand.NewSource(currentTime)
	newRand := rand.New(source)

	for i := range d {
		newPosition := newRand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
