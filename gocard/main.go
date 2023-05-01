package main

import (
	"gocard/deck"
)

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}
