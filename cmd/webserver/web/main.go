package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

type config struct {
	port          int
	staticDirPath string
}

func main() {
	cfg := config{
		port:          4000,
		staticDirPath: "./cmd/webserver/ui/static",
	}

	flag.IntVar(&cfg.port, "p", cfg.port, "HTTP port address")
	flag.StringVar(&cfg.staticDirPath, "sdp", cfg.staticDirPath, "HTTP port address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDirPath))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /snippet/view/{id}", handlers.SnippetView)
	mux.HandleFunc("GET /snippet/create", handlers.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", handlers.SnippetCreatePost)

	port := fmt.Sprintf(":%d", cfg.port)

	logger.Info("starting server", "port", port)
	logger.Error(http.ListenAndServe(port, mux).Error())
	os.Exit(1)
}
