package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
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
		data := templatedata.TemplateData{
			Snippets: snippets,
		}

		app.Render(w, r, http.StatusOK, "home", data)
	}
}
