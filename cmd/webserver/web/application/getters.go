package application

import "html/template"

func (app *Application) GetTemplateCache() map[string]*template.Template {
	return app.TemplateCache
}
