package main

import "fmt"

/*
	Criating a struct type is kinda creating an object to store data.
	An struct type can store varibles and even other structures.
*/
type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	// Creating a new person called p
	p := person{
		firstName: "Peter",
		lastName:  "Previatti",
		contact: contact{
			email:   "pvlpreviatti@gmail.com",
			zipCode: 19191919,
		},
	}

	p.updateName("Pedro")
	p.print()
}

// Simple function to print all of person's 'variables'
func (p person) print() {
	fmt.Printf("\nFirst Name: %v", p.firstName)
	fmt.Printf("\nLast Name: %v", p.lastName)
	fmt.Printf("\nEmail: %v", p.contact.email)
	fmt.Printf("\nZip Code: %v", p.contact.zipCode)
}

//	pointer function accepted when passing a 'person' or a pointer to a 'person'
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
