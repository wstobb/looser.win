package game

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Session struct {
	UUID    uuid.UUID
	Game    *Game
	Expires time.Time
}

func NewSession(w http.ResponseWriter, logger *zerolog.Logger) *Session {
	session := &Session{
		UUID:    uuid.New(),
		Game:    newGame(),
		Expires: time.Now().Add(time.Hour * 6),
	}

	session.SetCookie(w)

	logger.Debug().Str("uuid", session.UUID.String()).Msg("created new session")

	return session
}

func (s *Session) Reset(w http.ResponseWriter) {
	s.Game = newGame()
	s.Expires = time.Now().Add(time.Hour * 6)
	s.SetCookie(w)
}

func (s *Session) Start() {
	s.Game.Start()
}

func (s *Session) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    s.UUID.String(),
		Expires:  s.Expires,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}
