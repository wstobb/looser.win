package game

import (
	"math/rand"
)

type Deck struct {
	Cards []*Card
	Used  []*Card
}

func newDeck() *Deck {
	deck := Deck{
		Cards: make([]*Card, 0, 52),
		Used:  make([]*Card, 0),
	}
	for i := range 4 {
		for j := range 13 {
			deck.Cards = append(deck.Cards, newCard(i+1, j+1))
		}
	}
	return &deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Draw() *Card {
	if len(d.Cards) == 0 {
		return nil
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) Use(card *Card) {
	d.Used = append(d.Used, card)
}
