package main

import "fmt"


type englishBot struct{}
type spanishBot struct{}

type bot interface {
	greeting() string
}

func (b englishBot) greeting() string{
	return "Hello World"
}

func (b spanishBot) greeting() string{
	return "Hola Worlda"
}

func printGreeting(b bot){
	fmt.Println(b.greeting())
}

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()
	en := englishBot{}
	es := spanishBot{}
	printGreeting(en)
	printGreeting(es)
}