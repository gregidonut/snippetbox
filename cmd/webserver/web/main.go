package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

const (
	DEFAULT_PORT = ":4000"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./cmd/webserver/ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /snippet/view/{id}", handlers.SnippetView)
	mux.HandleFunc("GET /snippet/create", handlers.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", handlers.SnippetCreatePost)

	fmt.Printf("listening at %s\n", DEFAULT_PORT)
	log.Fatal(http.ListenAndServe(DEFAULT_PORT, mux))
}
