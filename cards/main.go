package main

func main() {
	cards := deck{"Ace of Diamons", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
