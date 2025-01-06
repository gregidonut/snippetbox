package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
)

func snippetView(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Info("called snippetView")
		defer app.Info("completeed snippetView handler logic")
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id <= 0 {
			http.NotFound(w, r)
			return
		}
		app.Info("snippetView successfully parsed from url", slog.Int("id", id))
		snippet, err := app.Get(id)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		b, err := json.MarshalIndent(snippet, "", "    ")
		if err != nil {
			app.ServerError(w, r, err)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(b)
	}
}
