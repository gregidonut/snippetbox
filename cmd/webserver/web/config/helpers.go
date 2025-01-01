package config

import (
	"log/slog"
	"net/http"
)

func (app *Application) ServerError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.Logger.Error(err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
	)
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func (app *Application) ClientError(
	w http.ResponseWriter,
	status int,
) {
	http.Error(
		w,
		http.StatusText(status),
		status,
	)
}
