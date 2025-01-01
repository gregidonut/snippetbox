package config

import (
	"log/slog"
	"os"
)

type Application struct {
	Logger *slog.Logger
}

func NewApplication() *Application {
	return &Application{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
	}
}
