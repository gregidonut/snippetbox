package config

import (
	"fmt"
	"log/slog"
	"os"
)

type RuntimeCFG struct {
	Port          int
	StaticDirPath string
}

func (c *RuntimeCFG) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func newRuntimeCFG() *RuntimeCFG {
	return &RuntimeCFG{
		Port:          4000,
		StaticDirPath: "./cmd/webserver/ui/static",
	}
}

type Application struct {
	Logger *slog.Logger
	*RuntimeCFG
}

func NewApplication() *Application {
	return &Application{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})),
		RuntimeCFG: newRuntimeCFG(),
	}
}
