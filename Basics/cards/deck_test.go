package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckLength := len(d)
	if deckLength != 52 {
		t.Errorf("Expected deck length of 52, but got %v", deckLength)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[deckLength-1] != "King of Hearts" {
		t.Errorf("Expected first card to be King of Hearts, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	deckLength := len(loadedDeck)
	if deckLength != 52 {
		t.Errorf("Expected deck length of 52, but got %v", deckLength)
	}

	os.Remove(filename)
}
