package game

import "github.com/wstobb/looser.win/internal/tools"

type Game struct {
	Rounds int
	Deck   *Deck
	Hands  []*Hand
}

func newGame() *Game {
	return &Game{
		Rounds: 0,
		Deck:   newDeck(),
		Hands:  tools.MakeSlice(4, newHand),
	}
}
