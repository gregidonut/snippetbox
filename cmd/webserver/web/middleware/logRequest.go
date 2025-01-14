package middleware

import (
	"log/slog"
	"net/http"

	"github.com/justinas/alice"
)

type App interface {
	Info(msg string, args ...any)
}

func LogRequest(app App) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ip     = r.RemoteAddr
				proto  = r.Proto
				method = r.Method
				uri    = r.URL.RequestURI()
			)
			app.Info("received request",
				slog.String("ip", ip),
				slog.String("proto", proto),
				slog.String("method", method),
				slog.String("uri", uri),
			)
			next.ServeHTTP(w, r)
		})
	}
}
