package game

type Inventory struct {
	Hand     []*Card
	Chips    int
	Discards int
	Folded   bool
}

func NewInventory() *Inventory {
	return &Inventory{
		Hand:     make([]*Card, 0, 5),
		Chips:    1000,
		Discards: 0,
		Folded:   false,
	}
}
