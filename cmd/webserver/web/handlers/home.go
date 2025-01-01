package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r
		w.Header().Add("Server", "Go")
		w.Header().Add("Content-Type", "text/html")

		files := []string{
			"./cmd/webserver/ui/html/base.tmpl.html",
			"./cmd/webserver/ui/html/partials/nav.tmpl.html",
			"./cmd/webserver/ui/html/pages/home.tmpl.html",
		}

		ts, err := template.ParseFiles(
			files...,
		)
		if err != nil {
			app.Logger.Error(err.Error(),
				slog.String("method", r.Method),
				slog.String("uri", r.URL.RequestURI()),
			)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
			app.Logger.Error(err.Error(),
				slog.String("method", r.Method),
				slog.String("uri", r.URL.RequestURI()),
			)
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
		}
	}
}
