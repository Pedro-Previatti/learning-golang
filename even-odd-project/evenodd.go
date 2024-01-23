package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type ins_s is an integer slice
type int_s []int

// function to return a slice of size i with random numbers that go up to s
func newNumbers(i int, s int) int_s {
	// creating seed to ensure ramdomization
	seed := rand.NewSource(time.Now().UnixNano())

	// making a new slice of type int_s and size i
	slice := make(int_s, i)

	// randomizing slice up to s
	slice.random(s, seed)

	return slice
}

// function to print a slice of int_s
func (s int_s) print() {
	for i := range s {
		fmt.Println(i, ": ", s[i])
	}
}

// function will return a randomized int_s slice
func (s int_s) random(i int, seed rand.Source) int_s {
	// randomizing from seed
	r := rand.New(seed)
	// testing random
	// fmt.Println(seed)

	// iterating through the slice and generationg random number
	for n := range s {
		s[n] = r.Intn(i) + 1
	}

	// testing random
	// s.print()

	return s
}

// simple function checking if generated numbers are even or odd
func (s int_s) checkEvenOdd() {
	for i := range s {
		if s[i]/2 == 0 {
			fmt.Printf("\n(%v: %2v): Even", i, s[i])
		} else {
			fmt.Printf("\n(%v: %2v): Odd", i, s[i])
		}
	}
}
