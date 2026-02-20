package server

import (
	"io/fs"
	"net/http"
)

func (s *Server) fileServers() {
	sub, err := fs.Sub(s.static, "static")
	if err != nil {
		panic(err)
	}

	s.mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(sub)))
}
