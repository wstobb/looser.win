package server

import (
	"net/http"
)

func (s *Server) indexHander(w http.ResponseWriter, r *http.Request) {
	s.pageCreator(w, "index.html")
}
