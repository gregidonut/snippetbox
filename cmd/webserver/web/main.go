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
	app, err := config.NewApplication()
	if err != nil {
		app.Error(err.Error(), slog.String("constructor", "config.NewApplication()"))
		os.Exit(1)
	}
	defer app.Close()

	configFilePath := flag.String("conf", config.DEFAULT_CONFIG_PATH, "specify a config file path")
	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "static directory path")
	flag.StringVar(&app.ConnStr, "cs", app.ConnStr, "postgresql connection string")
	flag.StringVar(&app.HtmlPagesTemplatePathGlob, "htp", app.HtmlPagesTemplatePathGlob, "pattern to match glob for templates should prob be *.tmpl.html")
	flag.Parse()

	if *configFilePath != config.DEFAULT_CONFIG_PATH {
		rcfg, err := config.NewRuntimeCFG(*configFilePath)
		if err != nil {
			app.Error(err.Error(), slog.String("constructor", "config.NewRuntimeCFG()"))
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
