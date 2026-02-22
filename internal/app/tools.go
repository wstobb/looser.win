package app

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/wstobb/looser.win/internal/game"
)

func (a *App) pageCreator(w http.ResponseWriter, name string) {
	page := template.Must(template.ParseFS(a.templates, filepath.Join("templates/pages", name)))
	if err := page.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *App) ensureSession(w http.ResponseWriter, r *http.Request) *game.Session {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		return game.NewSession(w, a.logger)
	}

	session := a.sessions[sessionCookie.Value]

	if session == nil {
		return game.NewSession(w, a.logger)
	}

	return session
}
