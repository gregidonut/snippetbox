package appconfig

import (
	"net/http"
	"time"

	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
)

type TemplateData struct {
	models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}

func (app *Application) NewTemplateData(r *http.Request) TemplateData {
	return TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
