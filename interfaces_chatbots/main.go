package main

import (
	"fmt"
)

// Interfaces helps to the reuse of code
// Interfaces are not generic types
// Interfaces implementation are implicit
// Interface definition
// Interfaces can contain other interfaces
type bot interface {
	getGreeting() string // function requiered to implement the interface
	// function inside interface can use multiple argument types and return types
	// getGreetings(typeA, tupeB)           (typeX, TypeY)
	//             list of argument types   list of return types

	// The type that want to implement it has to implement every functions inside the interface
}

// also called concrete type
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// if you are a type in this program with a function called 'getGreetings', and return a string, you are now honorary member of type bot
// implicit interface implementation
func (englishBot) getGreeting() string {
	return "Hi there!"
}

// implicit interface implementation
func (spanishBot) getGreeting() string {
	return "Hola amigos!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Overloading functions is not supported
// func printGreeting(sb spanishBot) {
// fmt.Println(sb.getGreeting())
// }
