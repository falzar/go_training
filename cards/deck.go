package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//  The new type will be a slice of Strings
// You can think is a type that extends from slice type
type deck []string

// type decoration
// Helps create a new "Instance"
func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "four"}

	// When there is a variable that we will not use but we need to skip, use "_"
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Create new deck from file. We cannot override a function changing arguments.
func newDeckFromFile(fileName string) deck {

	// Read file from disk
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		// exit program with status error 1
		os.Exit(1)
	}

	// Transform the readed data into first a slice of strings and then to a deck
	return deck(strings.Split(string(bs), ","))
}

// Receiver function
// d is a reference of this type, in analogy is like "this" or "self"
// by go convention, the name of receiver is a single letter witch is the starting letter of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Go can return multiple values from one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// type conversion type_desired(type_to_convert)
func (d deck) toString() string {

	// Join uses the second parameter to join elementes in the slice in the first argument
	return strings.Join([]string(d), ",")
}

// Returns error on error
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {

	// Creating a true Random number isntead default pseudorandom
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		rn := r.Intn(len(d) - 1)

		// swap elements
		d[i], d[rn] = d[rn], d[i]
	}
}
