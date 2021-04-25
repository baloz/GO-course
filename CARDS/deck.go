package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string //a slice of strings

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}
func (d deck) shuffle(){
	
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d{
		newPos := r.Intn(len(d) - 1)

		d[index], d[newPos] = d[newPos], d[index]
	}
}