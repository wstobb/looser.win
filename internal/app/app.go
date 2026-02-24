package app

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/wstobb/looser.win/internal/config"
	"github.com/wstobb/looser.win/internal/core"
	"github.com/wstobb/looser.win/internal/logging"
)

type App struct {
	logger    *zerolog.Logger
	mux       *chi.Mux
	static    fs.FS
	templates fs.FS
	sessions  map[string]*core.Session
}

func New(static, templates fs.FS) *App {
	logger := logging.New(zerolog.DebugLevel)
	return &App{
		logger:    logger,
		mux:       chi.NewMux(),
		static:    static,
		templates: templates,
		sessions:  make(map[string]*core.Session),
	}
}

func (a *App) Start() error {
	a.fileServers()
	a.routes()

	port := config.GetPort()

	a.logger.Info().Msg(fmt.Sprintf("starting http server at %s", port))
	if err := http.ListenAndServe(port, a.mux); err != nil {
		return err
	}

	return nil
}
