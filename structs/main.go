package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	email     string
	zipCode   int
}

func main() {
	p := person{
		firstName: "Pedro",
		lastName:  "Previatti",
		email:     "pvlpreviatti@gmail.com",
		zipCode:   19191919,
	}

	p.print()
}

func (p person) print() {
	fmt.Printf("\nFirst Name: %v", p.firstName)
	fmt.Printf("\nLast Name: %v", p.lastName)
	fmt.Printf("\nEmail: %v", p.email)
	fmt.Printf("\nZip Code: %v", p.zipCode)
}
