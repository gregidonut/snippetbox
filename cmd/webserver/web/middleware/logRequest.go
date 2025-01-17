package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appinterface"
	"github.com/justinas/alice"
)

func LogRequest(app appinterface.App) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			app.Info("received request",
				slog.String("ip", r.RemoteAddr),
				slog.String("proto", r.Proto),
				slog.String("method", r.Method),
				slog.String("uri", r.URL.RequestURI()),
			)
			next.ServeHTTP(w, r)
		})
	}
}
