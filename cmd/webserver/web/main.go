package main

import (
	"crypto/tls"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/application"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"
)

func main() {
	app, dbCloserFunc, err := application.NewApplication()
	if err != nil {
		app.Error(err.Error(), slog.String("constructor", "application.NewApplication()"))
		os.Exit(1)
	}
	if dbCloserFunc != nil {
		defer dbCloserFunc()
	}

	applicationFilePath := flag.String("conf", config.DEFAULT_CONFIG_PATH, "specify a application file path")
	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "static directory path")
	flag.StringVar(&app.ConnStr, "cs", app.ConnStr, "postgresql connection string")
	flag.StringVar(&app.HtmlTemplateDirPath, "htdp", app.HtmlTemplateDirPath, "dir path for the html templates")
	flag.StringVar(&app.TLSCertPath, "tlscp", app.TLSCertPath, "file path for the self signed tls cert")
	flag.StringVar(&app.TLSKeyPath, "tlskp", app.TLSKeyPath, "file path for the self signed tls key")
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
		TLSConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		},
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Info("starting server", slog.Int("port", app.Port))
	app.Error(
		srv.ListenAndServeTLS(app.TLSCertPath, app.TLSKeyPath).Error(),
	)
	os.Exit(1)
}
