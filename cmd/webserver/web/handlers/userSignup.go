package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func userSignup(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userSignup handler...")
		defer app.Debug("finished running userSignup handler")
		fmt.Fprintf(w, "Authenticate and log the user in...")
	}
}
