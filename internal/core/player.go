package core

type Player interface {
	Get() Player
	Draw()
	Discard()
	GetInventory() Inventory
}

type NPC interface {
	Player
	MakeDecision()
}
