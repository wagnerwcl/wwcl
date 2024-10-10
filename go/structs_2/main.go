package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Wagner")
	jim.print()
}

func (p *person) updateName(newFirstName string) { // We have a reciever and a argument.
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
