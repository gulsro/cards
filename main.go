package main

//import "fmt"

func main(){
	cards := newDeck()

	//cards.print()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	// cards := newDeckFromFile("jello")
	// cards.print()
	cards.shuffle()
	cards.print()
}