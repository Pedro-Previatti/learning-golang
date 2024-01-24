package main

import "fmt"

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
	p := person{
		firstName: "Peter",
		lastName:  "Previatti",
		contact: contact{
			email:   "pvlpreviatti@gmail.com",
			zipCode: 19191919,
		},
	}

	ptr := &p
	ptr.updateName("Pedro")
	p.print()
}

func (p person) print() {
	fmt.Printf("\nFirst Name: %v", p.firstName)
	fmt.Printf("\nLast Name: %v", p.lastName)
	fmt.Printf("\nEmail: %v", p.contact.email)
	fmt.Printf("\nZip Code: %v", p.contact.zipCode)
}

func (ptr *person) updateName(newFirstName string) {
	ptr.firstName = newFirstName
}
