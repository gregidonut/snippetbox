package application

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/form/v4"
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
	app.Error(err.Error(),
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

func (app *Application) DecodePostForm(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	app.Debug(fmt.Sprintf("r.PostForm: %#v", r.PostForm))

	if err := app.Decode(dst, r.PostForm); err != nil {
		var InvalidDecoderErr *form.InvalidDecoderError
		if errors.As(err, &InvalidDecoderErr) {
			panic(err)
		}
		return err
	}
	return nil
}
