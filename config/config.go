package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

type Config struct {
	App            AppConfig
	Server         ServerConfig
	Database       DatabaseConfig
	Auth           AuthConfig
	StorageSession StorageSessionConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Description string
	Authors     []string
}

type ServerConfig struct {
	Port              uint16
	CaseSensitive     bool
	EnablePrintRoutes bool
}

type DatabaseConfig struct {
	Host        string
	Port        uint16
	User        string
	Password    string
	Name        string
	SSLMode     string
	AutoMigrate bool
}

type AuthConfig struct {
	Cost              uint8
	SessionExpiration int
}

type StorageSessionConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

func GetConfig() *Config {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config
}
