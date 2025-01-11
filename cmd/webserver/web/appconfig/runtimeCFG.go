package appconfig

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

const DEFAULT_CONFIG_PATH = "./defaultConfig.yaml"

type RuntimeCFG struct {
	Port                int    `yaml:"Port"`
	StaticDirPath       string `yaml:"StaticDirPath"`
	ConnStr             string `yaml:"ConnStr"`
	HtmlTemplateDirPath string `yaml:"HtmlTemplateDirPath"`
}

func (c *RuntimeCFG) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func NewRuntimeCFG(app *Application, appconfigFilePath string) (*RuntimeCFG, error) {
	app.Debug("creating appconfiguration", slog.String("constructor", "NewRuntimeCFG"))
	defer app.Debug("finished creating appconfiguration", slog.String("constructor", "NewRuntimeCFG"))
	if appconfigFilePath != DEFAULT_CONFIG_PATH {
		// TODO: implement custom appconfig file logic
		return &RuntimeCFG{}, errors.New("appconfig file specified but not implemented yet")
	}
	yamlFile, err := os.ReadFile(appconfigFilePath)
	if err != nil {
		return nil, err
	}

	payload := RuntimeCFG{}

	err = yaml.Unmarshal(yamlFile, &payload)
	if err != nil {
		return nil, err
	}

	app.Info(fmt.Sprintf("current appconfig %#v", payload))

	if err = validate(payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

// validateStruct checks for zero values in the struct and returns an error with the field name
func validate(s interface{}) error {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsZero() {
			fieldName := t.Field(i).Name
			return fmt.Errorf("field '%s' has a zero value", fieldName)
		}
	}
	return nil
}
