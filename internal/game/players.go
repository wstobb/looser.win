package game

type Player interface {
	Draw(deck *Deck)
	Discard(index int, deck *Deck)
	Fold()
	Bet(amount int) (int, error)
	GetSeat() *Seat
	IsActive() bool
}

func newPlayers() []Player {
	players := make([]Player, 0, 4)
	players = append(players, newUser())
	players = append(players, newSquirrel())
	players = append(players, newCobson())
	players = append(players, newChud())

	return players
}
