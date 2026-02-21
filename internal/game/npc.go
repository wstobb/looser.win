package game

type NPC struct {
	*Inventory
}

func NewNPC() *NPC {
	return &NPC{
		Inventory: NewInventory(),
	}
}
