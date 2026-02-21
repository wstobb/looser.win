package app

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func (a *App) pageCreator(w http.ResponseWriter, name string) {
	page := template.Must(template.ParseFS(a.templates, filepath.Join("templates/pages", name)))
	if err := page.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
