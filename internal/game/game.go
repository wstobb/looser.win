package game

import (
	"fmt"

	"github.com/wstobb/looser.win/internal/tools"
)

type Game struct {
	Rounds  int
	Deck    *Deck
	Players []Player
}

func newGame() *Game {
	return &Game{
		Rounds:  0,
		Deck:    NewDeck(),
		Players: NewPlayers(),
	}
}

func (g *Game) Start() {
	g.Deck.Shuffle()
	g.Deal()

	for i, player := range g.Players {
		for j, card := range player.GetSeat().Hand {
			fmt.Printf("player %d | card %d: %d, %d\n", i, j, card.Suit, card.Rank)
		}
	}
}

func (g *Game) Deal() {
	for range 5 {
		for i := range len(g.Players) {
			var card *Card
			card, g.Deck.Cards = tools.Pop(g.Deck.Cards)
			g.Players[i].GetSeat().Hand = append(g.Players[i].GetSeat().Hand, card)
		}
	}
}
