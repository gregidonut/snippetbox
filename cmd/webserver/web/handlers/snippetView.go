package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "hello from snippetView page %d\n", id)
}
