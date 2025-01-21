package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		app.Error(err.Error(), slog.String("constructor", "application.NewApplication()"))
		os.Exit(1)
	}
	defer app.Close()

	applicationFilePath := flag.String("conf", config.DEFAULT_CONFIG_PATH, "specify a application file path")
	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "static directory path")
	flag.StringVar(&app.ConnStr, "cs", app.ConnStr, "postgresql connection string")
	flag.StringVar(&app.HtmlTemplateDirPath, "htdp", app.HtmlTemplateDirPath, "dir path for the html templates")
	flag.Parse()

	if *applicationFilePath != config.DEFAULT_CONFIG_PATH {
		rcfg, err := config.NewRuntimeCFG(app, *applicationFilePath)
		if err != nil {
			app.Error(err.Error(), slog.String("constructor", "application.NewRuntimeCFG()"))
			os.Exit(1)
		}
		app.RuntimeCFG = rcfg
	}

	srv := &http.Server{
		Addr:     app.GetPort(),
		Handler:  handlers.Routes(app),
		ErrorLog: slog.NewLogLogger(app.Handler(), slog.LevelError),
	}

	app.Info("starting server", slog.Int("port", app.Port))
	app.Error(
		srv.ListenAndServe().Error(),
	)
	os.Exit(1)
}
