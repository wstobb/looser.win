package app

import (
	"fmt"
	"net/http"
)

func (a *App) indexHander(w http.ResponseWriter, r *http.Request) {
	session := a.ensureSession(w, r)
	a.sessions[session.UUID.String()] = session
	a.logger.Debug().Msg(fmt.Sprintf("session %s loaded", session.UUID.String()))
	a.pageCreator(w, "index.html")
}
