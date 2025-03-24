package config

import (
	"ams-service/middlewares"
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

var CONFIG_LOG_PREFIX = "config.go"

type Config struct {
	ServerPort string
	Database   DatabaseConfig
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	var configFile string
	if runtime.GOOS == "windows" {
		configFile = "C:/DEV/db-config/local-config.yaml"
	} else {
		configFile = "/path/to/db-config/local-config.yaml"
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Loading configuration from: %s", CONFIG_LOG_PREFIX, configFile))
	middlewares.LogInfo(fmt.Sprintf("%s - Operating System: %s", CONFIG_LOG_PREFIX, runtime.GOOS))
	middlewares.LogInfo(fmt.Sprintf("%s - Architecture: %s", CONFIG_LOG_PREFIX, runtime.GOARCH))

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Configuration loaded successfully: %s", CONFIG_LOG_PREFIX, configFile))

	return &cfg, nil
}
