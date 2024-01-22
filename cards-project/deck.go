package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Creating type deck as a slice of strings
Similar to an array, the difference of a slice is that it is
dynamically-sized, flexible view into the elements of an array
*/
type deck []string

// function created to return a deck of cards
func newDeck() deck {
	// created variable cards of type deck
	cards := deck{}

	/*
	 created 2 slices of type string:
	 	cardSuits has all the possible suits of a deck of cards
	 	cardValues has all the possible values of a deck of cards
	*/
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// 2 for loops to go through the suits and values appending to the cards slice variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// function returns the cards slice variables
	return cards
}

// function receives a deck slice and prints it
func (d deck) print() {
	fmt.Println("\ndeck:")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function receives a deck of cards and returns a new random deck
func (d deck) shuffle() {
	// creating a seed every shuffle to make sure it is different every time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// array goes through the deck
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// this statement changes the positions of the cards in the array
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// this function gets a deck and a handsize as arguments
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// this is a helper function to the saveToFile() function, and it turns a deck into a string
func (d deck) toString() string {
	// this statement will join all the deck slice with a comma in between
	return strings.Join([]string(d), ",")
}

// this function creates a new file with the deck string
func (d deck) saveToFile(s string) error {
	return os.WriteFile(s, []byte(d.toString()), 0666)
}

// function to save a deck in a file
func newDeckFromFile(s string) deck {
	bs, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	sts := strings.Split(string(bs), ",")
	return deck(sts)
}
