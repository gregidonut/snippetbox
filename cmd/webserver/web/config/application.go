package config

import (
	"database/sql"
	"log/slog"
	"os"
)

type Application struct {
	*slog.Logger
	*RuntimeCFG
	*sql.DB
}

func NewApplication() (*Application, error) {
	// payload needs to always be returned regardles of error since
	// entire app depends on the existence of logger
	payload := &Application{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
	}

	rcfg, err := NewRuntimeCFG(DEFAULT_CONFIG_PATH)
	if err != nil {
		return payload, err
	}
	payload.RuntimeCFG = rcfg

	payload.checkDefaultConfigPathExists()

	db, err := payload.openDB()
	if err != nil {
		return payload, err
	}
	payload.DB = db

	return payload, nil
}

func (app *Application) checkDefaultConfigPathExists() {
	app.Debug("checking if default config file exists...")
	defer app.Debug("completed existence check")

	_, err := os.Stat(DEFAULT_CONFIG_PATH)
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

func (app Application) openDB() (*sql.DB, error) {
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
