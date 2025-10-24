package appservice

import (
	"github.com/amirtavakolian/quiz-game/pkg/configloader"
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

	cfgLoader := app.ConfigLoader.SetDelimiter(".").SetYamlpath(filepath.Join(yamlConfigPath, "config", "app.yaml")).Build()

	return cfgLoader.String("sms-provider"), os.Getenv("KAVENEGAR_API")
}
