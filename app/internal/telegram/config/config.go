package config

import (
	"gopkg.in/hedzr/errors.v3"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ApiToken string `yaml:"apiToken"`
}

func New() (*Config, error) {
	//filePath := os.Getenv("TG_CONFIG_PATH")
	filePath := "app/internal/telegram/config.yml"
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("failed to read config file: %v", err)
	}
	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.New("failed to unmarshall config yaml:%v", err)
	}
	return &config, nil
}
