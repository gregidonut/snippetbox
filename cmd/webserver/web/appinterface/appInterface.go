package appinterface

import (
	"context"
	"html/template"
	"net/http"
)

type App interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	ServerError(
		w http.ResponseWriter,
		r *http.Request,
		err error,
	)
	PopString(ctx context.Context, key string) string
	GetTemplateCache() map[string]*template.Template
}
