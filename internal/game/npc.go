package game

type NPC struct {
	*Inventory
}

func NewNPC() *NPC {
	return &NPC{
		Inventory: NewInventory(),
	}
}

func (n *NPC) Draw(deck *Deck)               {}
func (n *NPC) Discard(index int, deck *Deck) {}
func (n *NPC) Fold()                         {}

func (n *NPC) Bet(amount int) (int, error) {
	return 0, nil
}

func (n *NPC) Get() *Inventory {
	return n.Inventory
}

func (n *NPC) IsActive() bool {
	return true
}
