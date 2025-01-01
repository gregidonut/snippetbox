package handlers

import "net/http"

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	_ = r

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("save a new snippet"))
}
