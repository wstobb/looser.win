package app

import "time"

func (a *App) PruneSessionsService() func() {
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				for id, session := range a.sessions {
					if time.Now().After(session.Expires) {
						delete(a.sessions, id)
					}
				}
			}
		}
	}()

	return func() {
		done <- true
	}
}
