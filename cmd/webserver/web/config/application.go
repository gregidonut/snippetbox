package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

const DEFAULT_CONFIG_PATH = "./defaultConfig.yaml"

type RuntimeCFG struct {
	Port          int    `yaml:"Port"`
	StaticDirPath string `yaml:"StaticDirPath"`
	ConnStr       string `yaml:"ConnStr"`
}

func (c *RuntimeCFG) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func NewRuntimeCFG(configFilePath string) (*RuntimeCFG, error) {
	if configFilePath != DEFAULT_CONFIG_PATH {
		// TODO: implement custom config file logic
		return &RuntimeCFG{}, errors.New("config file specified but not implemented yet")
	}
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	payload := RuntimeCFG{}

	err = yaml.Unmarshal(yamlFile, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

type Application struct {
	*slog.Logger
	*RuntimeCFG
}

func NewApplication() (*Application, error) {
	// mandatoryLogger will segfault entire app if not returned with the application struct even
	// in error situations since we rely on the error logs that won't exist from main
	//
	mandatoryLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	rcfg, err := NewRuntimeCFG(DEFAULT_CONFIG_PATH)
	if err != nil {
		return &Application{
			Logger: mandatoryLogger,
		}, err
	}
	return &Application{
		Logger:     mandatoryLogger,
		RuntimeCFG: rcfg,
	}, nil
}
