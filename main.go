package main

func main() {
	cards := newDeck()
	hand, remain := deal(cards, 4)
	remain.print()
	hand.print()
}
