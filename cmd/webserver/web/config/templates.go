package config

import (
	"html/template"
	"path/filepath"
	"strings"
)

func (app *Application) NewTemplateCache() (map[string]*template.Template, error) {
	payload := map[string]*template.Template{}

	pages, err := filepath.Glob(app.HtmlPagesTemplatePathGlob)
	if err != nil {
		return nil, err
	}

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
