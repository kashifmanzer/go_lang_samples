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

func main() {
	kashif := person{
		firstName: "Kashif",
		lastName:  "Manzer",
		contactInfo: contactInfo{
			email:   "kmanzer3@gmail.com",
			zipCode: 800014,
		},
	}
	fmt.Printf("%+v 111111 ", kashif)
	kashifPointer := &kashif
	kashifPointer.updateName("Afaque")
	kashif.print()
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
