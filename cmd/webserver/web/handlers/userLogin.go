package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func userLogin(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userLogin handler...")
		defer app.Debug("finished running userLogin handler")
		fmt.Fprintf(w, "Display a form for signing up a new user...")
	}
}
