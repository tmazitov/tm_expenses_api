package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
)

type config struct {
	DB postgresql.Config
}

func getenvDefault(variable string, defaultValue string) string {
	value := os.Getenv(variable)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func NewConfig() (*config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("config: failed to load .env: %w", err)
	}

	dbPort, err := strconv.Atoi(getenvDefault("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("config: DB_PORT must be a number: %w", err)
	}

	return &config{
		DB: postgresql.Config{
			Host:     getenvDefault("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getenvDefault("DB_USER", "expense_client"),
			Password: getenvDefault("DB_PASSWORD", "expense_client"),
			DBName:   getenvDefault("DB_NAME", "expense_db"),
			SSLMode:  getenvDefault("DB_SSL_MODE", "disable"),
		},
	}, nil
}
