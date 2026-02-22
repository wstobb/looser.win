package game

type Squirrel struct {
	*Seat
	Number int
}

func newSquirrel() *Squirrel {
	return &Squirrel{
		Seat:   newSeat(),
		Number: 1,
	}
}

func (s *Squirrel) Draw(deck *Deck) {}

func (s *Squirrel) Discard(index int, deck *Deck) {}

func (s *Squirrel) Fold() {}

func (s *Squirrel) Bet(amount int) (int, error) {
	return 0, nil
}

func (s *Squirrel) GetSeat() *Seat {
	return s.Seat
}

func (s *Squirrel) IsActive() bool {
	return !s.Folded
}

func (s *Squirrel) DecideAction(g *Game) {}
