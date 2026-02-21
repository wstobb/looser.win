package game

import (
	"fmt"
)

type Game struct {
	Rounds  int
	Deck    *Deck
	Players []Player
}

func newGame() *Game {
	return &Game{
		Rounds:  0,
		Deck:    newDeck(),
		Players: newPlayers(),
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
			card := g.Deck.Draw()
			g.Players[i].GetSeat().Hand = append(g.Players[i].GetSeat().Hand, card)
		}
	}
}
