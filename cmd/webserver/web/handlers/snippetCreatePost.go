package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

		expires, err := strconv.Atoi(r.PostForm.Get("expires"))
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		form := templatedata.NewSnippetCreateFormData(
			r.PostForm.Get("title"),
			r.PostForm.Get("content"),
			expires,
		)

		if len(form.FieldErrors) > 0 {
			data := app.NewTemplateData(r)
			data.SnippetCreateFormData = *form
			app.Render(w, r, http.StatusUnprocessableEntity, "create", data)
			return
		}

		id, err := app.Insert(*form)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
