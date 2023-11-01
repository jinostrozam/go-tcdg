package main

func main() {
	// cards := newDeck()
	// fmt.Println("\n--Deck--")
	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("\n--Hand--")
	// hand.print()

	// fmt.Println("\n--Remaining deck--")
	// remainingCards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
