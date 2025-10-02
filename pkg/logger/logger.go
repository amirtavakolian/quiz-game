package logger

import (
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/configloader"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

const (
	developMode    = "development-mode"
	productionMode = "production-mode"
)

type Logger struct {
	ConfigLoader configloader.ConfigLoader
}

func New() Logger {
	return Logger{ConfigLoader: *configloader.NewConfigLoader()}
}

func (l Logger) Log() *zap.Logger {
	var config zap.Config
	configPath, err := os.Getwd()

	if err != nil {
		panic(err.Error())
	}

	cfgLoader := l.ConfigLoader.SetPrefix("APP_").SetYamlpath(filepath.Join(configPath, "pkg", "logger", "config.yaml")).SetDelimiter(".").SetDivider("_").Build()

	if cfgLoader.String("current.mode") == developMode {
		err := cfgLoader.Unmarshal(developMode, &config)

		if err != nil {
			panic(err.Error())
		}
	} else {
		cfgLoader.Unmarshal(productionMode, &config)
	}

	logger, err := config.Build()

	if err != nil {
		panic(err)
	}

	return logger
}
