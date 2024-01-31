package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()

	// cards.toString()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()
}
