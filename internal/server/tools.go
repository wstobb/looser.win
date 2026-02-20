package server

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func (s *Server) pageCreator(w http.ResponseWriter, name string) {
	page := template.Must(template.ParseFS(s.templates, filepath.Join("templates/pages", name)))
	if err := page.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
