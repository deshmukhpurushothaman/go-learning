package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex.firstName = "Alex"
	// alex.lastName = "Anerson"
	// fmt.Println(alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Ken",
		contact: contactInfo{
			email:   "test@example.com",
			zipCode: "123456",
		},
	}
	// jimPointer := &jim // Converts value to memory address
	jim.updateName("Jimmy")
	jim.print()
}

// *person converts address to value
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
