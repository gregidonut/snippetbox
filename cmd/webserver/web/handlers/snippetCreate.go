package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func snippetCreate(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("ran home handler")
		defer app.Info("finished running home handler")
		data := templatedata.NewTemplateData[*templatedata.SnippetCreateFormData](r, app)
		render(app, w, r, http.StatusOK, "create", data)
	}
}
