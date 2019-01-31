package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first array item to be Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected first array item to be King of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()

	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand size of 5, but got hand size of %v", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected remaining deck size of 47, but got %v", len(remainingDeck))
	}
}
