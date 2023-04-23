package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	DBName   string `yaml:"DBName"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func New() (*Config, error) {
	filePath := "config.yml"
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("cannot find config file")
	}
	return &Config{
		DBName:   viper.GetString("database.DBName"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.User"),
		Password: viper.GetString("database.password"),
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
