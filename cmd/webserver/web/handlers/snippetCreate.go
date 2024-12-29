package handlers

import "net/http"

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Write([]byte("hello from snippetcreate page"))
}
