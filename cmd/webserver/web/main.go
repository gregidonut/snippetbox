package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

func main() {
	app := config.NewApplication()

	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "HTTP port address")
	flag.Parse()

	app.Logger.Info("starting server", slog.Int("port", app.Port))
	app.Logger.Error(
		http.ListenAndServe(app.GetPort(), handlers.Routes(app)).Error(),
	)
	os.Exit(1)
}
