package main

import "fmt"

// type bot is an interface
type bot interface {
	// when we declared this, the interface is basically saying that if you are a type in this package with a function getGreeting() that returns a string, than now you are an honorary member of type bot
	getGreeting() string
}

// two types that do the same thing but with a different language
type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

/*
	Both the englishBot and the Portuguese bot need a function to print the greeting.
	Since repeating the function without a receiver won't work, we have to create a new type bot.
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (portugueseBot) getGreeting() string {
	return "Salve"
}
