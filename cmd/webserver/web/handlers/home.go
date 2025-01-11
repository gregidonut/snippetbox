package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appconfig"
)

func home(app *appconfig.Application) http.HandlerFunc {
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
		data := appconfig.TemplateData{
			Snippets: snippets,
		}

		app.Render(w, r, http.StatusOK, "home", data)
	}
}
