package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func userLogin(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userLogin handler...")
		defer app.Debug("finished running userLogin handler")
		data := templatedata.New[*templatedata.UserLoginFormData](r, app)
		render(app, w, r, http.StatusOK, "login", data)
	}
}
