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
	htmlPagesTemplatePathGlob := filepath.Join(app.HtmlTemplateDirPath, "pages/*.tmpl.html")
	app.Debug(
		fmt.Sprintf("HtmlPagesTemplatePathGlob: %s", htmlPagesTemplatePathGlob),
		slog.String("constructor", "NewTemplateCache"),
	)
	pages, err := filepath.Glob(htmlPagesTemplatePathGlob)
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

		ts, err := template.ParseFiles(
			filepath.Join(app.HtmlTemplateDirPath, "base.tmpl.html"),
		)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(
			filepath.Join(app.HtmlTemplateDirPath, "partials/*.tmpl.html"),
		)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		payload[name] = ts
	}

	return payload, nil
}
