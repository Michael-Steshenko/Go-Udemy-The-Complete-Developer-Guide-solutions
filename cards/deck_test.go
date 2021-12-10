package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	myDeck := newDeck()

	if len(myDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(myDeck))
	}

	if myDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", myDeck[0])
	}

	if myDeck[len(myDeck)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", myDeck[len(myDeck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckt")

	deck := newDeck()
	deck.saveToFile("_deckt")

	loadedDeck := newDeckFromFile("_deckt")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deckt")

}
