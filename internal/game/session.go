package game

import (
	"fmt"

	"github.com/google/uuid"
)

type Session struct {
	UUID uuid.UUID
	Game *Game
}

func NewSession() *Session {
	return &Session{
		UUID: uuid.New(),
		Game: newGame(),
	}
}

func (s *Session) Reset() {
	s.Game = newGame()
}

func (s *Session) Start() {
	fmt.Printf("session: %s\n", s.UUID.String())
	s.Game.Start()
}
