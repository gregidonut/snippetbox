package appconfig

import "github.com/gregidonut/snippetbox/cmd/webserver/internal/models"

type TemplateData struct {
	models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}
