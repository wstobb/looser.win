package app

import (
	"io/fs"
	"net/http"
)

func (a *App) fileServers() {
	sub, err := fs.Sub(a.static, "static")
	if err != nil {
		panic(err)
	}

	a.mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(sub)))
}
