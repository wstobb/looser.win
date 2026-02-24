package core

const (
	Hearts = iota + 1
	Diamonds
	Clubs
	Spades
)

const (
	Ace = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit int
	Rank int
}
