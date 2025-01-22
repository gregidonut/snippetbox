package handlers

import (
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/middleware"
	"github.com/justinas/alice"
)

func Routes(app *application.Application) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(app.StaticDirPath))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(home(app)))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(snippetView(app)))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(snippetCreate(app)))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(snippetCreatePost(app)))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(userSignup(app)))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(userSignupPost(app)))
	mux.Handle("GET /user/login", dynamic.ThenFunc(userLogin(app)))
	mux.Handle("POST /user/login", dynamic.ThenFunc(userLoginPost(app)))
	mux.Handle("POST /user/logout", dynamic.ThenFunc(userLogoutPost(app)))

	standard := alice.New(
		middleware.RecoverPanic(app),
		middleware.LogRequest(app),
		middleware.CommonHeaders,
	)
	return standard.Then(mux)
}
