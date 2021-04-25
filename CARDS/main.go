package main
import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	//cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println()
	fmt.Println()
	remainingCards.print()
}
