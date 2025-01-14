package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appconfig"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/middleware"
	"github.com/justinas/alice"
)

func Routes(app *appconfig.Application) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.StaticDirPath))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home(app))
	mux.HandleFunc("GET /snippet/view/{id}", snippetView(app))
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost(app))

	standard := alice.New(
		middleware.LogRequest(app),
		middleware.CommonHeaders,
	)
	return standard.Then(mux)
}
