package app

import (
	"io/fs"

	"github.com/rs/zerolog"
	"github.com/wstobb/looser.win/internal/logging"
	"github.com/wstobb/looser.win/internal/server"
)

type App struct {
	logger *zerolog.Logger
	server *server.Server
}

func New(static, templates fs.FS) *App {
	logger := logging.New(zerolog.DebugLevel)
	return &App{
		server: server.New(static, templates),
		logger: logger,
	}
}

func (a *App) Start() error {
	if err := a.server.Start(a.logger); err != nil {
		return err
	}
	return nil
}
