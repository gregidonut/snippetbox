package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func userLoginPost(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userLoginPost handler...")
		defer app.Debug("finished running userLoginPost handler")
		fmt.Fprintf(w, "Authenticate and log the user in...")
	}
}
