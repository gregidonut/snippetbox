package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appconfig"
)

func snippetCreatePost(app *appconfig.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("called snippetCreate handler")
		defer app.Debug("finished snippetCreate handler")

		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
		expires := 7

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
