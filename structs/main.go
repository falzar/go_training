package main

import (
	"fmt"
)

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
	// Can be ssigned by position or defining attribute name like next line
	john := person{firstName: "John", lastName: "Tangarife"}
	fmt.Println(john)

	//Default values for struct elements
	var vito person
	fmt.Println(vito)
	fmt.Printf("%+v\n", vito)

	// Can be assigned by element
	vito.firstName = "Juan"
	vito.lastName = "Urrango"

	fmt.Printf("%+v\n", vito)

	esteban := person{
		firstName: "Esteban",
		lastName:  "Alarcon",
		contact: contactInfo{
			email:   "nabetse@some.co",
			zipCode: 123456, //for multistructs definition, every single line in the definition should have a comma even if is the last line
		},
	}

	fmt.Printf("%+v\n", esteban)

	// & give us the access to the memory address of esteban person struct
	estebanPointer := &esteban

	fmt.Println("pointer:", estebanPointer)

	// work the same if a pass a value or a pointer.
	vito.updateName("Carlos")
	estebanPointer.updateName("Alejandro")

	esteban.print()
	vito.print()
}

// Go is a pass by value language. p is a new copy on memory of original caller.
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// * give access to memory address
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// Use pointers with: int, float, string, bool, struct.
// Dont need to use pointers with: slice, map, channel, pointer, funtion
