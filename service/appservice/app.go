package appservice

import (
	"github.com/amirtavakolian/adapter-and-repository-pattern-in-golang/pkg/configloader"
	"os"
	"path/filepath"
)

type AppService struct {
	ConfigLoader *configloader.ConfigLoader
}

func NewAppService(configLoader *configloader.ConfigLoader) AppService {
	return AppService{ConfigLoader: configLoader}
}

func (app AppService) GetSmsProvider() (string, string) {
	yamlConfigPath, _ := os.Getwd()

	cfgLoader := app.ConfigLoader.SetDelimiter(".").SetYamlpath(filepath.Join(yamlConfigPath, "config", "app.yaml")).
		SetPrefix("APP_").SetDivider("_").Build()

	return cfgLoader.String("sms-provider"), cfgLoader.String("kavenegar.api")
}
