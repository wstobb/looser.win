package core

import "math/rand"

type Deck struct {
	Active    []*Card
	Discarded []*Card
}

func NewDeck() *Deck {
	// Allocate deck
	deck := &Deck{
		Active:    make([]*Card, 0, 52),
		Discarded: make([]*Card, 0, 52),
	}

	// Populate cards
	for i := range 4 {
		for j := range 13 {
			deck.Active = append(deck.Active, &Card{Suit: i, Rank: j})
		}
	}

	return deck
}

func (d *Deck) WithJoker(amount int) *Deck {
	for range amount {
		d.Active = append(d.Active, &Card{Suit: 0, Rank: 14})
	}

	return d
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Active), func(i, j int) {
		d.Active[i], d.Active[j] = d.Active[j], d.Active[i]
	})
}

func (d *Deck) Draw() *Card {
	if len(d.Active) == 0 {
		return nil
	}

	card := d.Active[0]
	d.Active = d.Active[1:]
	return card
}

func (d *Deck) Discard(card *Card) {
	d.Discarded = append(d.Discarded, card)
}
