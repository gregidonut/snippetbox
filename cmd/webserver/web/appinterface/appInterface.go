package appinterface

import "net/http"

type App interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	ServerError(
		w http.ResponseWriter,
		r *http.Request,
		err error,
	)
}
