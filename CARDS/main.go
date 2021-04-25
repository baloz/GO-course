package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of Spades", "King of spades"}
	cards = append(cards, "Queen of Hearts")

	cards.print()
}
