package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Bob",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateFirstName("Jimmy")
	jim.print()
}
