package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func userSignupPost(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userSignupPost handler...")
		defer app.Debug("finished running userSignupPost handler")

		form := &templatedata.UserSignupFormData{}

		if err := app.DecodePostForm(r, &form); err != nil {
			app.Error(err.Error())
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		form.Validate()
		if !form.Valid() {
			data := templatedata.NewTemplateData[*templatedata.UserSignupFormData](r, app)
			data.Form = form
			render(app, w, r, http.StatusUnprocessableEntity, "signup", data)
			return
		}

		fmt.Fprintf(w, "Create new user....")
	}
}
