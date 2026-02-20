package game

type Hand struct {
	Cards []*Card
}

func newHand() *Hand {
	return &Hand{
		Cards: make([]*Card, 0),
	}
}
