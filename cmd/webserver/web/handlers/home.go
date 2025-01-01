package handlers

import (
	"html/template"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
)

func home(app *config.Application) http.HandlerFunc {
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
			app.ServerError(w, r, err)
			return
		}

		if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
			app.ServerError(w, r, err)
		}
	}
}
