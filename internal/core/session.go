package core

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID      string
	Game    Game
	Active  bool
	Expires time.Time
}

func NewSession(w http.ResponseWriter) *Session {
	session := &Session{
		ID:      uuid.NewString(),
		Active:  false,
		Expires: time.Now().Add(time.Hour * 6),
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    session.ID,
		Expires:  session.Expires,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	return session
}

func (s *Session) Start() {
	s.Active = true
}
