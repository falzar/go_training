package main

import "fmt"

//  The new type will be a slice of Strings
// You can think is a type that extends from slice type
type deck []string

// Receiver function
// d is a reference of this type, in analogy is like "this" or "self"
// by go convention, the name of receiver is a single letter witch is the starting letter of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
