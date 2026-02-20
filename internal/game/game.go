package game

type Game struct {
	Rounds int
	Deck   *Deck
	Hands  []*Hand
}
