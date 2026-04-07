package config

import (
	"strings"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewCORSConfig() cors.Config {
	// Get allowed origins from environment variable (comma-separated)
	// Example: "http://localhost:8082,http://127.0.0.1:8082,http://192.168.1.100:8082"
	allowedOriginsRaw := getenvDefault("ALLOWED_ORIGINS", "http://localhost:8082")
	allowedOrigins := strings.Split(allowedOriginsRaw, ",")

	// Trim whitespace from origins
	for i := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
	}

	return cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}
}
