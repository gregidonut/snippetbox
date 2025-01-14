package middleware

import (
	"fmt"
	"net/http"

	"github.com/justinas/alice"
)

func RecoverPanic(app App) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Add("Connection", "close")
					app.ServerError(w, r, fmt.Errorf("%s", err))
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
