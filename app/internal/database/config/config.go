package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
)

// Config - struct of database configurations
type Config struct {
	DBName   string `yaml:"DBName"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// New creates new db config
func New() (*Config, error) {
	filePath := os.Getenv("db_config_path")
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
}
