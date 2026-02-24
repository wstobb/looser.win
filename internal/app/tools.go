package app

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/wstobb/looser.win/internal/core"
)

func (a *App) pageCreator(w http.ResponseWriter, name string) {
	page := template.Must(template.ParseFS(a.templates, filepath.Join("templates/pages", name)))
	if err := page.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *App) ensureSession(w http.ResponseWriter, r *http.Request) *core.Session {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		return core.NewSession(w)
	}

	session := a.sessions[sessionCookie.Value]

	if session == nil {
		return core.NewSession(w)
	}

	return session
}
