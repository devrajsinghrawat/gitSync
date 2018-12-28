package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected Deck length 16, but received %v", len(d))
	}

	if d[0] != "Ace_Of_Spades" {
		t.Errorf("Expected 1st card in Deck is Ace_Of_Spades, but received %v", d[0])
	}

	if d[len(d)-1] != "Four_Of_Clubs" {
		t.Errorf("Expected Last card in Deck is Four_Of_Clubs, but received %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewSDeckFromFile(t *testing.T) {
	d := newDeck()

	os.Remove("_decktesting")

	err := d.saveToFile("_decktesting")

	if err != nil {
		t.Errorf("Expected Deck to save at disc, but received error %v", err)
	}

	dr := newDeckFromFile("_decktesting")
	if len(dr) != 16 {
		t.Errorf("Expected Deck length 16, but received %v", len(dr))
	}
	os.Remove("_decktesting")
}
