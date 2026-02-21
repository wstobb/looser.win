package app

import (
	"net/http"

	"github.com/wstobb/looser.win/internal/game"
)

func (a *App) indexHander(w http.ResponseWriter, r *http.Request) {
	game.NewSession(w, a.logger)
	a.pageCreator(w, "index.html")
}
