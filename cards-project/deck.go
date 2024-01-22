package main

import "fmt"

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
	cardSuits := []string{"Spaces", "Clubs", "Hearts", "Diamonds"}
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
	for i, card := range d {
		fmt.Println(i, card)
	}
}
