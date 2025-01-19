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

		if err := r.ParseForm(); err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}
		app.Debug(fmt.Sprintf("r.PostForm: %#v", r.PostForm))

		form := templatedata.SnippetCreateFormData{}

		if err := app.Decode(&form, r.PostForm); err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}
		form.Validate()

		if !form.Valid() {
			data := app.NewTemplateData(r)
			data.SnippetCreateFormData = form
			app.Render(w, r, http.StatusUnprocessableEntity, "create", data)
			return
		}

		id, err := app.Insert(form)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
