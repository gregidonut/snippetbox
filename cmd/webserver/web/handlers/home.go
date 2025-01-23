package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func home(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("ran home handler")
		defer app.Info("finished running home handler")
		snippets, err := app.Latest()
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		data := templatedata.NewTemplateData[*templatedata.BlankFormData](r, app)
		data.Snippets = snippets

		render(app, w, r, http.StatusOK, "home", data)
	}
}
