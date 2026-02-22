package game

type Seat struct {
	Hand     []*Card
	Chips    int
	Discards int
	Folded   bool
}

func newSeat() *Seat {
	return &Seat{
		Hand:     make([]*Card, 0, 5),
		Chips:    1000,
		Discards: 0,
		Folded:   false,
	}
}
