package handlers

import (
	"html/template"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
)

func home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("ran home handler")
		defer app.Info("finished running home handler")
		snippets, err := app.Latest()
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		w.Header().Add("Server", "Go")
		w.Header().Add("Content-Type", "text/html")

		files := []string{
			"./cmd/webserver/ui/html/base.tmpl.html",
			"./cmd/webserver/ui/html/partials/nav.tmpl.html",
			"./cmd/webserver/ui/html/pages/home.tmpl.html",
		}

		data := templateData{
			Snippets: snippets,
		}

		ts, err := template.ParseFiles(
			files...,
		)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		if err = ts.ExecuteTemplate(w, "base", data); err != nil {
			app.ServerError(w, r, err)
		}
	}
}
