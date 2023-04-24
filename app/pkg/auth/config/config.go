package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

func New() (Config, error) {
	filePath := "config.yml"
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, errors.New("cannot find config file")
	}
	return Config{
		Host: viper.GetString("sessions.host"),
		Port: viper.GetString("sessions.port"),
	}, nil
}
