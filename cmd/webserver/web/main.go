package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/appconfig"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

func main() {
	app, err := appconfig.NewApplication()
	if err != nil {
		app.Error(err.Error(), slog.String("constructor", "appconfig.NewApplication()"))
		os.Exit(1)
	}
	defer app.Close()

	appconfigFilePath := flag.String("conf", appconfig.DEFAULT_CONFIG_PATH, "specify a appconfig file path")
	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "static directory path")
	flag.StringVar(&app.ConnStr, "cs", app.ConnStr, "postgresql connection string")
	flag.StringVar(&app.HtmlTemplateDirPath, "htdp", app.HtmlTemplateDirPath, "dir path for the html templates")
	flag.Parse()

	if *appconfigFilePath != appconfig.DEFAULT_CONFIG_PATH {
		rcfg, err := appconfig.NewRuntimeCFG(app, *appconfigFilePath)
		if err != nil {
			app.Error(err.Error(), slog.String("constructor", "appconfig.NewRuntimeCFG()"))
			os.Exit(1)
		}
		app.RuntimeCFG = rcfg
	}

	app.Info("starting server", slog.Int("port", app.Port))
	app.Error(
		http.ListenAndServe(app.GetPort(), handlers.Routes(app)).Error(),
	)
	os.Exit(1)
}
