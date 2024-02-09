package config

import (
	"github.com/spf13/viper"
	"traive-engineering-challenge/internal/support"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault(support.DatabaseURL, "postgresql://postgres:password@localhost:5432/transactions-app_development?sslmode=disable")

	var config Config

	config.DatabaseURL = viper.GetString(support.DatabaseURL)

	return &config, nil
}
