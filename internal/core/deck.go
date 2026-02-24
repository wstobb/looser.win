package core

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

func (d *Deck) Draw() *Card {
	if len(d.Active) == 0 {
		return nil
	}

	card := d.Active[0]
	d.Active = d.Active[1:]
	return card
}
