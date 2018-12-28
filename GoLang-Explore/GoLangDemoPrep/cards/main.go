package main

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards.saveToFile("myDeck.txt")
	cards := newDeckFromFile("myDeck.txt")
	cards.shuffle()
	cards.print()
}
