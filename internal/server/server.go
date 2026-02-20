package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/wstobb/looser.win/internal/config"
	"github.com/wstobb/looser.win/internal/logging"
)

type Server struct {
	mux       *chi.Mux
	static    fs.FS
	templates fs.FS
}

func New(static, templates fs.FS, logger *zerolog.Logger) *Server {
	mux := chi.NewMux()
	mux.Use(logging.LoggerMiddleware(logger))
	return &Server{
		mux:       mux,
		static:    static,
		templates: templates,
	}
}

func (s *Server) Start(logger *zerolog.Logger) error {
	s.fileServers()
	s.routes()

	logger.Info().Msg(fmt.Sprintf("starting http server at %s", config.GetPort()))
	if err := http.ListenAndServe(config.GetPort(), s.mux); err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}
