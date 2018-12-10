package main

func main() {
	cards := newDeck()
	cards.shuffleDeck()
	cards.print()

	//cards.saveToFile("mycards.txt")
	// cards := newDeckFromFile("mycards.txt")

}
