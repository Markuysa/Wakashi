package config

import (
	"github.com/spf13/viper"
	"gopkg.in/hedzr/errors.v3"
	"os"
)

type Config struct {
	ApiToken string `yaml:"apiToken"`
}

// New creates new config object
func New() (*Config, error) {
	filePath := os.Getenv("tg_config_path")
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("cannot find config file")
	}
	return &Config{
		ApiToken: viper.GetString("telegram.apiToken"),
	}, nil
	//yamlFile, err := os.ReadFile(filePath)
	//if err != nil {
	//	return nil, errors.New("failed to read config file: %v", err)
	//}
	//config := make(map[string]Config)
	//
	//err = yaml.Unmarshal(yamlFile, &config)
	//if err != nil {
	//	return nil, errors.New("failed to unmarshall config yaml:%v", err)
	//}
	//return config, nil
}
