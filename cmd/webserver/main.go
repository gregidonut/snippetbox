package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	DEFAULT_PORT = ":4000"
)

func home(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Write([]byte("hello from home"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("hello from snippetView page %d\n", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Write([]byte("hello from snippetcreate page"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	fmt.Printf("listening at %s\n", DEFAULT_PORT)
	log.Fatal(http.ListenAndServe(DEFAULT_PORT, mux))
}
