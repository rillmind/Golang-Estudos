package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	expect := 52

	if len(deck) != expect {
		t.Errorf("Expected deck length of %v, but got %v", expect, len(deck))
	}

	if deck[0] != "Ace of Hearts" {
		t.Errorf("Expected 'Ace of Hearts', but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Spades" {
		t.Errorf("Expected 'King of Spades', but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	filename := "_deckTesting"

	os.Remove(filename)

	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	expect := 52

	if len(loadedDeck) != expect {
		t.Errorf("Expected deck length of %v, but got %v", expect, len(loadedDeck))
	}

	os.Remove(filename)
}
