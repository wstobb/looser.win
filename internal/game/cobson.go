package game

type Cobson struct {
	*Seat
	Number int
}

func newCobson() *Cobson {
	return &Cobson{
		Seat:   newSeat(),
		Number: 2,
	}
}

func (c *Cobson) Draw(deck *Deck) {}

func (c *Cobson) Discard(index int, deck *Deck) {}

func (c *Cobson) Fold() {}

func (c *Cobson) Bet(amount int) (int, error) {
	return 0, nil
}

func (c *Cobson) GetSeat() *Seat {
	return c.Seat
}

func (c *Cobson) IsActive() bool {
	return !c.Folded
}

func (c *Cobson) DecideAction(g *Game) {}
