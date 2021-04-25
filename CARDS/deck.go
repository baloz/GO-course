package main

import "fmt"

type deck []string //a slice of strings

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string {"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards  = append(cards, value +" of "+suit)
		}
	}
	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}