package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
)

func home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("ran home handler")
		defer app.Info("finished running home handler")

		w.Header().Add("Server", "Go")
		// w.Header().Add("Content-Type", "text/html")
		//
		// files := []string{
		// 	"./cmd/webserver/ui/html/base.tmpl.html",
		// 	"./cmd/webserver/ui/html/partials/nav.tmpl.html",
		// 	"./cmd/webserver/ui/html/pages/home.tmpl.html",
		// }
		//
		// ts, err := template.ParseFiles(
		// 	files...,
		// )
		// if err != nil {
		// 	app.ServerError(w, r, err)
		// 	return
		// }
		//
		// if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
		// 	app.ServerError(w, r, err)
		// }
		snippets, err := app.Latest()
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		b, err := json.MarshalIndent(snippets, "", "    ")
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		w.Header().Add("Content-Type", "Application/json")
		w.Write(b)
	}
}
