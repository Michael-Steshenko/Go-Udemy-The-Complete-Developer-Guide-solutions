package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// creates a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal 'numCards' from deck 'd'
func (d *deck) deal(numCards int) deck {
	hand := (*d)[:numCards]
	*d = (*d)[numCards:]
	return hand
}

func (d deck) toString() string {
	// a deck is really a slice of type string,
	// so we don't actually need to convert deck to a slice of type string.
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	deckString := strings.Split(string(byteSlice), ",")
	return deck(deckString)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
