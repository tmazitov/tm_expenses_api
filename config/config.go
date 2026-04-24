package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

type config struct {
	CORS        cors.Config
	DB          *DB
	GoogleOAuth *GoogleOAuth
	JWT         *JWT
	Cache       *Cache
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

	cors := NewCORSConfig()

	jwt, err := NewJWT()
	if err != nil {
		return nil, err
	}

	googleOAuth, err := NewGoogleOauth()
	if err != nil {
		return nil, err
	}

	cache, err := NewCache()
	if err != nil {
		return nil, err
	}

	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	return &config{
		DB:          db,
		GoogleOAuth: googleOAuth,
		CORS:        cors,
		JWT:         jwt,
		Cache:       cache,
	}, nil
}
