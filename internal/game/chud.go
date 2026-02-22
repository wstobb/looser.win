package game

type Chud struct {
	*Seat
	Number int
}

func newChud() *Chud {
	return &Chud{
		Seat:   newSeat(),
		Number: 3,
	}
}

func (c *Chud) Draw(deck *Deck) {}

func (c *Chud) Discard(index int, deck *Deck) {}

func (c *Chud) Fold() {}

func (c *Chud) Bet(amount int) (int, error) {
	return 0, nil
}

func (c *Chud) GetSeat() *Seat {
	return c.Seat
}

func (c *Chud) IsActive() bool {
	return !c.Folded
}

func (c *Chud) DecideAction(g *Game) {}
