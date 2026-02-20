package server

func (s *Server) routes() {
	s.mux.Get("/", s.indexHander)
}
