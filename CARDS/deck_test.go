package main

import "testing"

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