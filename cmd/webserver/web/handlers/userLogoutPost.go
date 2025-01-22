package handlers

import (
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
)

func userLogoutPost(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Debug("running userLogoutPost handler...")
		defer app.Debug("finished running userLogoutPost handler")
		fmt.Fprintf(w, "Logout the user...")
	}
}
