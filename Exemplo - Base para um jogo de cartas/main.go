package main

import "fmt"

func main() {
	newDeck := newDeck()

	// hand, remainingCards := deal(newDeck, 9)

	// hand.print()
	// remainingCards.print()

	// newDeck.saveToFile("myFirstDeck.txt")
	//

	newDeck.shuffle()

	fmt.Print(newDeck)
}
