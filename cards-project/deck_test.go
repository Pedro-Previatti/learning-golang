package main

import (
	"os"
	"reflect"
	"testing"
)

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

// this function will test if the shuffle function is working
func TestShuffle(t *testing.T) {
	d, s := newDeck(), newDeck()
	s.shuffle()

	// reflect will compare the slices to check if they are the same and it fails if they are
	if reflect.DeepEqual(d, s) == true {
		t.Errorf("Shuffle function not working")
	}
}

// this function will test if it is saving a deck to a new file file
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// making sure there is no file with this name at the start of the test
	os.Remove("_decktesting")

	s := 52
	d := newDeck()

	// creating file to test
	d.saveToFile("_decktesting")

	// loading file to test
	loadedDeck := newDeckFromFile("_decktesting")

	// if both the created and the loaded files have the same size it will pass
	if len(loadedDeck) != s {
		t.Errorf("Expected loaded deck from archive to be of length %v, but got length %v", s, len(loadedDeck))
	}

	// removing garbage
	os.Remove("_decktesting")
}
