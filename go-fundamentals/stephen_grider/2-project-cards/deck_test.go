package main

import (
	"os"
	"testing"
)

// t is test handler
func TestNewDeck(t *testing.T) {

	d := newDeck()

	// Code to make sure that a deck is created with x number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Code to make sure that the last card is a Four of Clubs
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove previous testing files (if any)
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
