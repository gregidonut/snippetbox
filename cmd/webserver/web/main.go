package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

type runtimeCFG struct {
	port          int
	staticDirPath string
}

func (c *runtimeCFG) getPort() string {
	return fmt.Sprintf(":%d", c.port)
}

func main() {
	app := config.NewApplication()
	cfg := runtimeCFG{
		port:          4000,
		staticDirPath: "./cmd/webserver/ui/static",
	}

	flag.IntVar(&cfg.port, "p", cfg.port, "HTTP port address")
	flag.StringVar(&cfg.staticDirPath, "sdp", cfg.staticDirPath, "HTTP port address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDirPath))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", handlers.Home(app))
	mux.HandleFunc("GET /snippet/view/{id}", handlers.SnippetView)
	mux.HandleFunc("GET /snippet/create", handlers.SnippetCreate)
	mux.HandleFunc("POST /snippet/create", handlers.SnippetCreatePost)

	app.Logger.Info("starting server", slog.Int("port", cfg.port))
	app.Logger.Error(
		http.ListenAndServe(cfg.getPort(), mux).Error(),
	)
	os.Exit(1)
}
