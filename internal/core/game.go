package core

type Game interface {
	Get() Game
	Start()
	Reset()
}
