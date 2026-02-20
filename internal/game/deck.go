package game

import "math/rand"

type Deck struct {
	Cards []*Card
}

func NewDeck() *Deck {
	deck := Deck{Cards: make([]*Card, 0, 52)}
	for i := range 4 {
		for j := range 13 {
			deck.Cards = append(deck.Cards, NewCard(i+1, j+1))
		}
	}
	return &deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
