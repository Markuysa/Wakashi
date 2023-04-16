package config

import (
	"gopkg.in/hedzr/errors.v3"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Host    string `yaml:"host"`
	Driver  string `yaml:"driver"`
	Port    string `yaml:"port"`
	SslMode string `yaml:"sslMode"`
	User    string `yaml:"user"`
}

func New() (*Config, error) {
	filePath := os.Getenv("T_CONFIG_PATH")
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("failed to read config file: %v", err)
	}
	var config *Config

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, errors.New("failed to unmarshall config yaml:%v", err)
	}
	return config, nil
}
