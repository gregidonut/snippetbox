package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Header().Add("Server", "Go")
	w.Write([]byte("hello from home"))
}
