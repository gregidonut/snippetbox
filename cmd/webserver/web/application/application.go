package application

import (
	"database/sql"
	"html/template"
	"log/slog"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/config"

	_ "github.com/lib/pq"
)

type Application struct {
	*slog.Logger
	*config.RuntimeCFG
	*models.SnippetModel
	TemplateCache map[string]*template.Template
	*form.Decoder
	*scs.SessionManager
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

	rcfg, err := config.NewRuntimeCFG(payload, config.DEFAULT_CONFIG_PATH)
	if err != nil {
		return payload, err
	}
	payload.RuntimeCFG = rcfg

	payload.checkDefaultConfigPathExists()

	db, err := payload.openDB()
	if err != nil {
		return payload, err
	}

	payload.SnippetModel = models.NewSnippetModel(db)

	tc, err := payload.NewTemplateCache()
	if err != nil {
		return payload, err
	}
	payload.TemplateCache = tc

	payload.Decoder = form.NewDecoder()

	payload.SessionManager = scs.New()
	payload.Store = postgresstore.New(db)
	payload.Lifetime = 12 * time.Hour

	return payload, nil
}

func (app *Application) checkDefaultConfigPathExists() {
	app.Debug("checking if default application file exists...")
	defer app.Debug("completed existence check")

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

func (app *Application) openDB() (*sql.DB, error) {
	app.Debug("started openDB func")
	defer app.Debug("finished openDB func")

	db, err := sql.Open("postgres", app.ConnStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
