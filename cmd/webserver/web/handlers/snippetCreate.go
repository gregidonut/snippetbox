package handlers

import "net/http"

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Write([]byte("hello from snippetcreate page"))
}
