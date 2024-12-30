package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Header().Add("Server", "Go")
	w.Header().Add("Content-Type", "text/html")

	files := []string{
		"./cmd/webserver/ui/html/base.tmpl.html",
		"./cmd/webserver/ui/html/partials/nav.tmpl.html",
		"./cmd/webserver/ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(
		files...,
	)
	if err != nil {

		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)
		return
	}

	if err = ts.ExecuteTemplate(w, "base", nil); err != nil {
		log.Print(err.Error())
		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)
	}
}
