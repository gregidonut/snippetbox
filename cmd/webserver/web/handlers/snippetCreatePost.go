package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func snippetCreatePost(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("called snippetCreate handler")
		defer app.Debug("finished snippetCreate handler")

		form := &templatedata.SnippetCreateFormData{}

		if err := app.DecodePostForm(r, &form); err != nil {
			app.Error(err.Error())
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		form.Validate()
		if !form.Valid() {
			data := templatedata.New[*templatedata.SnippetCreateFormData](r, app)
			data.Form = form
			render(app, w, r, http.StatusUnprocessableEntity, "create", data)
			return
		}

		id, err := app.SnippetModel.Insert(form)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		app.Put(r.Context(), "flash", "Snippet created successfully!")

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
