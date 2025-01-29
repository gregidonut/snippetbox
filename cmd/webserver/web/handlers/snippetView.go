package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func snippetView(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("called snippetView")
		defer app.Info("completed snippetView handler logic")
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id <= 0 {
			http.NotFound(w, r)
			return
		}
		app.Info("snippetView successfully parsed from url", slog.Int("id", id))
		snippet, err := app.SnippetModel.Get(id)
		if err != nil {
			if errors.Is(errors.Unwrap(err), models.ErrNoRecord) {
				http.NotFound(w, r)
				return
			}
			app.ServerError(w, r, err)
			return
		}
		app.Debug("parsed tmpl files")

		data := templatedata.New[*templatedata.BlankFormData](r, app)
		data.Snippet = snippet

		render(app, w, r, http.StatusOK, "view", data)
	}
}
