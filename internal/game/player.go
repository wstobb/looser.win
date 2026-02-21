package game

type Player interface {
	Draw(deck *Deck)
	Discard(index int, deck *Deck)
	Fold()
	Bet(amount int) (int, error)
	Get() *Inventory
	IsActive() bool
}

func NewPlayers() []Player {
	players := make([]Player, 0, 4)

	players = append(players, NewUser())

	for range 3 {
		players = append(players, NewNPC())
	}

	return players
}
