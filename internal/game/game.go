package game

import (
	"fmt"

	"github.com/wstobb/looser.win/internal/tools"
)

type Game struct {
	Rounds int
	Deck   *Deck
	Hands  []*Hand
}

func newGame() *Game {
	return &Game{
		Rounds: 0,
		Deck:   NewDeck(),
		Hands:  tools.MakeSlice(4, NewHand),
	}
}

func (g *Game) Start() {
	g.Deck.Shuffle()
	g.Deal()

	for i, hand := range g.Hands {
		for j, card := range hand.Cards {
			fmt.Printf("hand %d | card %d: %d, %d\n", i, j, card.Suit, card.Rank)
		}
	}
}

func (g *Game) Deal() {
	for range 5 {
		for i := range len(g.Hands) {
			var card *Card
			card, g.Deck.Cards = tools.Pop(g.Deck.Cards)
			g.Hands[i].Cards = append(g.Hands[i].Cards, card)
		}
	}
}
