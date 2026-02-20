package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/wstobb/looser.win/internal/logging"
)

type App struct {
	mux    *chi.Mux
	logger *zerolog.Logger
}

func New() *App {
	logger := logging.New(zerolog.DebugLevel)
	return &App{
		mux:    chi.NewMux(),
		logger: logger,
	}
}

func (a *App) Start() error {
	a.logger.Info().Msg("starting http server")
	if err := http.ListenAndServe(":80", a.mux); err != nil {
		return err
	}
	return nil
}
