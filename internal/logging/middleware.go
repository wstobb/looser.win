package logging

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			defer func() {
				logger.Info().
					Str("method", r.Method).
					Str("url", r.RequestURI).
					Str("remote_addr", r.RemoteAddr).
					Str("user_agent", r.UserAgent()).
					Int("status", ww.Status()).
					Int("bytes", ww.BytesWritten()).
					Dur("latency", time.Since(start)).
					Msg("request")
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
