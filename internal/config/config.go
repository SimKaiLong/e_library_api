package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"3000"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"host=localhost user=user password=pass dbname=lib sslmode=disable"`
	DBType      string `env:"DB_TYPE" envDefault:"memory"`
	Environment string `env:"APP_ENV" envDefault:"development"`
}

func LoadConfig() (*Config, error) {
	// Load the .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
