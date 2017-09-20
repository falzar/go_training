package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// Called automatically by testing framework. Argument t is a test handler.
// Use "Test" prefix to create test names
func TestNewDeck(t *testing.T) {

	// Arrange
	expectedLength := 16
	expectedFirstElement := "Ace of Spades"
	expectedLastElement := "Four of Clubs"

	// Act
	d := newDeck()

	// Assert
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}

	if d[0] != expectedFirstElement {
		t.Errorf("Expected %v at first element but got %v", expectedFirstElement, d[0])
	}

	if d[len(d)-1] != expectedLastElement {
		t.Errorf("Expected %v at last element but got %v", expectedLastElement, d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {

	// Arrange
	d := newDeck()
	testFileName := "_decktesting"

	// Act
	d.saveToFile(testFileName)
	_, err := ioutil.ReadFile(testFileName)

	// Assert
	if err != nil {
		t.Error(err)
	}

	// clean
	os.Remove(testFileName)
}

func TestNewDeckFromFile(t *testing.T) {

	// Arrange
	d := newDeck()
	testFileName := "_decktesting"
	d.saveToFile(testFileName)
	expectedLength := 16
	expectedFirstElement := "Ace of Spades"
	expectedLastElement := "Four of Clubs"

	// Act
	deckFromFile := newDeckFromFile(testFileName)

	// Assert
	if len(deckFromFile) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(deckFromFile))
	}

	if deckFromFile[0] != expectedFirstElement {
		t.Errorf("Expected %v at first element but got %v", expectedFirstElement, deckFromFile[0])
	}

	if deckFromFile[len(deckFromFile)-1] != expectedLastElement {
		t.Errorf("Expected %v at last element but got %v", expectedLastElement, deckFromFile[len(deckFromFile)-1])
	}

	// clean
	os.Remove(testFileName)
}
