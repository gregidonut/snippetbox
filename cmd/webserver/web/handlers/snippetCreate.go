package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func snippetCreate(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("ran home handler")
		defer app.Info("finished running home handler")
		data := app.NewTemplateData(r)
		app.Render(w, r, http.StatusOK, "create", data)
	}
}
