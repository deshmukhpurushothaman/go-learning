package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	cards.shuffle()
	cards.print()
	// Any variable with type deck can call print func
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
}
