package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetDefault("DATABASE_URL", "postgresql://postgres:password@localhost:5432/transactions-app_development?sslmode=disable")

	var config Config

	config.DatabaseURL = viper.GetString("DATABASE_URL")

	return &config, nil
}
