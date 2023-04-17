package config

import (
	"gopkg.in/hedzr/errors.v3"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DBName   string `yaml:"DBName"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func New() (map[string]Config, error) {
	filePath := "app/internal/database/config.yml"
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("failed to read config file: %v", err)
	}
	config := make(map[string]Config)

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errors.New("failed to unmarshall config yaml:%v", err)
	}
	return config, nil
}
