package app

func (a *App) routes() {
	a.mux.Get("/", a.indexHander)
}
