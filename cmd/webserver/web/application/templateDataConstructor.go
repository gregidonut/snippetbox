package application

import (
	"net/http"
	"time"

	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"
)

type TemplateData struct {
	models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
	Flash       string
	templatedata.SnippetCreateFormData
}

func (app *Application) NewTemplateData(r *http.Request) TemplateData {
	return TemplateData{
		CurrentYear: time.Now().Year(),
		Flash:       app.PopString(r.Context(), "flash"),
		SnippetCreateFormData: templatedata.SnippetCreateFormData{
			Expires: 7,
		},
	}
}
