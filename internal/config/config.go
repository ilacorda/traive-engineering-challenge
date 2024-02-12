package config

import (
	"github.com/spf13/viper"
	"traive-engineering-challenge/internal/support"
)

type Config struct {
	DatabaseURL string
}

// LoadConfig loads the application configuration from environment variables. It returns a Config struct and an error if the configuration could not be loaded.
// TODO [Improvements - application configs] env-specific configuration files (e.g. .yaml) could be used to load the configuration and override the default values.
// Besides, it is possible to integrate a secret manager (e.g. AWS Secrets Manager) to store sensitive information such as database credentials.
func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault(support.DatabaseURL, "postgresql://postgres:password@localhost:5432/transactions-app_development?sslmode=disable")

	var config Config

	config.DatabaseURL = viper.GetString(support.DatabaseURL)

	return &config, nil
}
