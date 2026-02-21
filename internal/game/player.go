package game

type Player interface {
	Draw(deck *Deck)
	Discard(index int, deck *Deck)
	Fold()
	Bet(amount int) (int, error)
	Get() *Inventory
	IsActive() bool
}

type NPC struct {
	*Inventory
}

func NewNPC() *NPC {
	return &NPC{
		Inventory: NewInventory(),
	}
}
