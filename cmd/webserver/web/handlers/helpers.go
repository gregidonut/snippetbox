package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appinterface"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

func render[T templatedata.FormData](
	app appinterface.App,
	w http.ResponseWriter,
	r *http.Request,
	status int,
	page string,
	data templatedata.TemplateData[T],
) {
	app.Debug(fmt.Sprintf("running render for %s", page))
	defer app.Debug(fmt.Sprintf("finished running render for %s", page))

	ts, ok := app.GetTemplateCache()[page]
	if !ok {
		pageNames := []string{}
		for k := range app.GetTemplateCache() {
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
