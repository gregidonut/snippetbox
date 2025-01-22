package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func userSignupPost(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userSignupPost handler...")
		defer app.Debug("finished running userSignupPost handler")
		fmt.Fprintf(w, "Create a new user...")
	}
}
