package config

import (
	"fmt"
	"strconv"
)

type DB struct {
	Host     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Port     int
}

func NewDB() (*DB, error) {

	dbPort, err := strconv.Atoi(getenvDefault("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("db config error : db port must be a number: %w", err)
	}

	return &DB{
		Host:     getenvDefault("DB_HOST", "localhost"),
		Port:     dbPort,
		User:     getenvDefault("DB_USER", "expense_client"),
		Password: getenvDefault("DB_PASSWORD", "expense_client"),
		DBName:   getenvDefault("DB_NAME", "expense_db"),
		SSLMode:  getenvDefault("DB_SSL_MODE", "disable"),
	}, nil
}
