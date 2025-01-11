package config

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
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

func (app *Application) Render(w http.ResponseWriter, r *http.Request, status int, page string, data templatedata.TemplateData) {
	app.Debug(fmt.Sprintf("running render for %s", page))
	defer app.Debug(fmt.Sprintf("finished running render for %s", page))

	ts, ok := app.TemplateCache[page]
	if !ok {
		pageNames := []string{}
		for k := range app.TemplateCache {
			pageNames = append(pageNames, k)
		}
		app.ServerError(w, r,
			fmt.Errorf("the template %s doesn't exist\n curr pages: %#v", page, pageNames),
		)
		return
	}

	buf := new(bytes.Buffer)

	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		app.ServerError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
