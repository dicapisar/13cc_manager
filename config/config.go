package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Description string
	Authors     []string
}

type ServerConfig struct {
	Port               uint16
	CaseSensitive      bool
	EnamblePrintRoutes bool
}

type DatabaseConfig struct {
	Host string
	Port uint16
	User string
	Pass string
	Name string
}

func GetConfig() *Config {
	var config *Config

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}
