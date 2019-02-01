package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type person2 struct {
	firstName string
	lastName  string
	contact   contactInfo
	// You do not need to specify a property name for a nested struct
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// This is not a great syntax to use due to the assignment happening in order
	// if you change the order of the assignment in the struct then all initialisations of the struct will now be wrong
	ryan := person{"Ryan", "Schranz"}
	fmt.Println(ryan)

	// This is a better way to use a struct as there assignment is explicit
	ryan2 := person{firstName: "Ryan", lastName: "Schranz"}
	fmt.Println(ryan2)

	// Create a variable of ryan3 as type person
	// This creates a zero value struct
	// This means the following data types are
	// string ""
	// int    0
	// float  0
	// bool   false
	var ryan3 person
	fmt.Println(ryan3)
	// Printf with %+v will print all properties of a struct
	fmt.Printf("%+v\n", ryan3)
	// To assign values to this struct initialisation
	ryan3.firstName = "Ryan"
	ryan3.lastName = "Schranz"
	fmt.Printf("%+v", ryan3)

	// Create a new person with their contact info
	ryan4 := person2{
		firstName: "Ryan",
		lastName:  "Schranz",
		contact: contactInfo{
			email:   "ryan@schranz.com",
			zipCode: 94000,
		},
		// If no property name if given for a struct, we need to use the struct name as the property name
		contactInfo: contactInfo{
			email:   "ryan.schranz@schranz.com",
			zipCode: 95000,
		},
	}
	// This will not update the firstName due to nothing being returned
	// This is because the pointer is not being updated
	ryan4.updateName("RyanS")

	// To make sure we update the correct point and memory section
	// We need to create a pointer reference for ryan4
	ryan4Pointer := &ryan4
	ryan4Pointer.updateName2("RyanS")
	ryan4.print()
}

// Receiver function for the struct of person2
func (p person2) print() {
	fmt.Printf("%+v", p)
}

func (p person2) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// To make a receiver function update the correct pointer it needs to be changed to the following
func (pointerToPerson2 *person2) updateName2(newFirstName string) {
	(*pointerToPerson2).firstName = newFirstName
}
