package game

type NPC interface {
	Player
	DecideAction(g *Game)
}
