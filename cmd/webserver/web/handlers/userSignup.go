package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func userSignup(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userSignup handler...")
		defer app.Debug("finished running userSignup handler")
		data := templatedata.NewTemplateData[*templatedata.UserSignupFormData](r, app)
		render(app, w, r, http.StatusOK, "signup", data)
	}
}
