package main

import "fmt"

type deck[] string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spade", "Diamond", "Hearts", "Club"}
	nums := []string{"1","2","3","4"}
	for _, suit := range suits {
		for _, num := range nums {
			cards = append(cards, suit + " " + num)
		}
	}
	return cards
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
