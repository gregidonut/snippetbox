package config

import (
	"errors"
	"fmt"
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
