package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"

	// card := "Ace of Spades"
	// card = newCard()

	// fmt.Println(card)

	//Slice
	cards := []string{"Ace of Spades", newCard()}

	// Adding elements to slice
	cards = append(cards, "Six of Hearts")

	// sub slices example cards[startIncludedIndex:endNotIncludedIndex]
	fmt.Println(cards[0:2])
	fmt.Println(cards[:2])
	fmt.Println(cards[0:])
	fmt.Println(cards[:])

	// Using deck
	// cards := deck{"Ace of Spades", newCard()}

	// Using deck instance
	deckOfCards := newDeck()

	// Using deck's receiver function
	fmt.Println("Original deck cards")
	deckOfCards.print()
	fmt.Printf("\n")

	// Usind deck's multiple return function
	hand, remainingCards := deal(deckOfCards, 5)

	fmt.Println("hand")
	hand.print()
	fmt.Printf("\n")
	fmt.Println("Remaining cards")
	remainingCards.print()
	fmt.Printf("\n")

	fmt.Println("Original deck cards to string:", deckOfCards.toString())
	fmt.Printf("\n")

	// IO Operations
	hand.saveToFile("my_cards") // Verify file is created

	storedHand := newDeckFromFile("my_cards")
	fmt.Println("Stored Hand:", storedHand.toString()) // Test file read

	// Random
	storedHand.shuffle()
	fmt.Println("Shuffled stored Hand:", storedHand.toString()) // Test file read

	fmt.Println(newDeckFromFile("mu_cards")) // Test error handling

}

// Function declaration
func newCard() string {
	return "Five of Diamonds"
}
