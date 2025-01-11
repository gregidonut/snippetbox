package config

import (
	"fmt"
	"html/template"
	"log/slog"
	"path/filepath"
	"strings"
)

func (app *Application) NewTemplateCache() (map[string]*template.Template, error) {
	app.Debug("creating template cache", slog.String("constructor", "NewTemplateCache"))
	defer app.Debug("finished creating template cache", slog.String("constructor", "NewTemplateCache"))

	payload := map[string]*template.Template{}

	app.Debug(
		fmt.Sprintf("HtmlPagesTemplatePathGlob: %s", app.HtmlPagesTemplatePathGlob),
		slog.String("constructor", "NewTemplateCache"),
	)
	pages, err := filepath.Glob(app.HtmlPagesTemplatePathGlob)
	if err != nil {
		return nil, err
	}

	app.Debug("looping through pages path", slog.String("constructor", "NewTemplateCache"))
	for _, page := range pages {
		var (
			baseFName = filepath.Base(page)
			tmplFName = strings.TrimSuffix(baseFName, filepath.Ext(baseFName))
			name      = strings.TrimSuffix(tmplFName, filepath.Ext(tmplFName))
		)

		files := []string{
			"./cmd/webserver/ui/html/base.tmpl.html",
			"./cmd/webserver/ui/html/partials/nav.tmpl.html",
			page,
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		payload[name] = ts
	}

	return payload, nil
}
