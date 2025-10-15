package configloader

import (
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
	"strings"
)

type ConfigLoader struct {
	Prefix    string
	Delimiter string
	YamlPath  string
	Divider   string
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (loader *ConfigLoader) SetPrefix(prefix string) *ConfigLoader {
	loader.Prefix = prefix
	return loader
}

func (loader *ConfigLoader) SetDelimiter(delimiter string) *ConfigLoader {
	loader.Delimiter = delimiter
	return loader
}

func (loader *ConfigLoader) SetYamlpath(yamlPath string) *ConfigLoader {
	loader.YamlPath = yamlPath
	return loader
}

func (loader *ConfigLoader) SetDivider(divider string) *ConfigLoader {
	loader.Divider = divider
	return loader
}

func (loader *ConfigLoader) Build() *koanf.Koanf {
	k := koanf.New(loader.Delimiter)

	if loader.YamlPath != "" {
		if err := k.Load(file.Provider(loader.YamlPath), yaml.Parser()); err != nil {
			panic(err.Error())
		}
	}

	if loader.Prefix != "" {
		if err := godotenv.Load(); err != nil {
			panic(err.Error())
		}

		if err := k.Load(env.Provider(loader.Prefix, loader.Delimiter, func(s string) string {
			return strings.Replace(strings.ToLower(
				strings.TrimPrefix(s, loader.Prefix)), loader.Divider, loader.Delimiter, -1)
		}), nil); err != nil {
			log.Fatalf("Error loading env: %v", err)
		}
	}

	return k
}
