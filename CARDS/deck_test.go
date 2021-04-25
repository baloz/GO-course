package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if(len(d) != 16){
		t.Errorf("Expected length 16 but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[15])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}