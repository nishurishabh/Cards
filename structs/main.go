package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Struct to represent a person
type person struct {
	firstName string
	lastname  string
	contact   contactInfo
}

func main() {
	rishabh := person{
		firstName: "Rishabh",
		lastname:  "Patel",
		contact: contactInfo{
			email:   "nishu@gmail.com",
			zipCode: 221005,
		},
	}
	rishabh.updateName("Nishu")
	rishabh.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// func main() {
// 	// ways of declaring structs
// 	// alex := person{"Alex", "Anderson"}     						  poor way because it takes ordering into consideration
// 	// alex := person{firstName: "Alex", lastname: "Anderson"}        Better way because we specify each field without need of order.

// 	var alex person // Initializes person with empty strings
// 	alex.firstName = "Rishabh"
// 	alex.lastname = "Patel"
// 	fmt.Println(alex)
// }
