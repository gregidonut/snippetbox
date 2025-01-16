package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appconfig"
)

func snippetCreatePost(app *appconfig.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("called snippetCreate handler")
		defer app.Debug("finished snippetCreate handler")

		if err := r.ParseForm(); err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		expires, err := strconv.Atoi(r.PostForm.Get("expires"))
		if title == "" || content == "" || err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		app.Debug(fmt.Sprintf("r.PostForm: %#v", r.PostForm))

		id, err := app.Insert(
			title,
			content,
			expires,
		)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
