package game

import (
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

func (s *Session) Start() {
	s.Game.start()
}
