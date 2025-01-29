package handlers

import (
	"errors"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
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
			data := templatedata.New[*templatedata.UserSignupFormData](r, app)
			data.Form = form
			render(app, w, r, http.StatusUnprocessableEntity, "signup", data)
			return
		}
		err := app.UserModel.Insert(form)
		if err != nil {
			if errors.Is(err, models.ErrDuplicateEmail) {
				form.AddFieldError("email", "Email address is already in use")

				data := templatedata.New[*templatedata.UserSignupFormData](r, app)
				data.Form = form

				render(app, w, r, http.StatusUnprocessableEntity, "signup", data)
			}
		}

		app.Put(r.Context(), "flash", "account successfully created!")

		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
	}
}
