package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/handlers"

	_ "github.com/lib/pq"
)

func checkDefaultConfigPathExists(app *config.Application) {
	_, err := os.Stat(config.DEFAULT_CONFIG_PATH)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		app.Warn("default filepath not exist make sure to specify it with -conf")
		return
	}
	app.Error(err.Error())
	os.Exit(1)
}

func main() {
	app, err := config.NewApplication()
	if err != nil {
		app.Error(err.Error(), slog.String("constructor", "config.NewApplication()"))
		os.Exit(1)
	}
	checkDefaultConfigPathExists(app)

	configFilePath := flag.String("conf", config.DEFAULT_CONFIG_PATH, "specify a config file path")
	flag.IntVar(&app.Port, "p", app.Port, "HTTP port address")
	flag.StringVar(&app.StaticDirPath, "sdp", app.StaticDirPath, "static directory path")
	flag.StringVar(&app.ConnStr, "cs", app.ConnStr, "postgresql connection string")
	flag.Parse()

	if *configFilePath != config.DEFAULT_CONFIG_PATH {
		rcfg, err := config.NewRuntimeCFG(*configFilePath)
		if err != nil {
			app.Error(err.Error(), slog.String("constructor", "config.NewRuntimeCFG()"))
			os.Exit(1)
		}
		app.RuntimeCFG = rcfg
	}

	db, err := openDB(app)
	if err != nil {
		app.Error("from openDB()", slog.String("err", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	app.Info("starting server", slog.Int("port", app.Port))
	app.Error(
		http.ListenAndServe(app.GetPort(), handlers.Routes(app)).Error(),
	)
	os.Exit(1)
}

func openDB(app *config.Application) (*sql.DB, error) {
	app.Debug("started openDB func")
	defer app.Debug("finished openDB func")

	db, err := sql.Open("postgres", app.ConnStr)
	if err == nil {
		return nil, err
	}

	err = db.Ping()
	if err == nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
