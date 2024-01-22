package main

import "testing"

// this function will test the size of a new deck created by the newDeck() function
func TestNewDeckSize(t *testing.T) {
	d := newDeck()
	s := 52

	// if the length of a new deck is any different than 52 it will fail
	if len(d) != s {
		t.Errorf("Expected deck length of %v, but got %v", s, len(d))
	}
}

// this function will test the first and last card from a new deck created by the newDeck() function
func TestNewDeckFirstAndLastCards(t *testing.T) {
	d := newDeck()
	// as seen in newDeck(), the first and last cards created in the deck should be as follows
	f := "Ace of Spades"
	l := "King of Diamonds"

	// if the first card isn't the 'Ace of Spades' it will fail
	if d[0] != f {
		t.Errorf("Expected first card to be %v, but got %v", f, d[0])
	}

	// if the last card isn't the 'King of Diamonds' it will fail
	if d[51] != l {
		t.Errorf("Expected last card to be %v, but got %v", l, d[51])
	}
}
