package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/tmazitov/ayda-order-service.git/api/docs"
	"github.com/tmazitov/ayda-order-service.git/api/rest"
	"github.com/tmazitov/ayda-order-service.git/api/rest/middleware"
	"github.com/tmazitov/ayda-order-service.git/config"
	"github.com/tmazitov/ayda-order-service.git/internal/app"
	infra "github.com/tmazitov/ayda-order-service.git/internal/infrastructure"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/google"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/jwt"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/redis"
	"github.com/tmazitov/ayda-order-service.git/pkg/validator"
)

// @title           Expense Tracker API
// @version         1.0
// @host            localhost:8080
func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatalf("service launch failed: %v", err)
	}

	infra, err := infra.NewInfrastructure(infra.InfrastructureParams{
		DBConfig: postgresql.Config{
			Host:     c.DB.Host,
			Port:     c.DB.Port,
			User:     c.DB.User,
			Password: c.DB.Password,
			SSLMode:  c.DB.SSLMode,
			DBName:   c.DB.DBName,
		},
		CacheParams: redis.CacheParams{
			Addr: c.Cache.Addr,
			DB:   c.Cache.DB,
		},
		GoogleOAuthParams: google.OAuthProviderParams{
			ClientId: c.GoogleOAuth.ClientId,
		},
		JwtParams: jwt.StorageParams{
			Secret:     c.JWT.Secret,
			AccessTTL:  c.JWT.AccessTTL,
			RefreshTTL: c.JWT.RefreshTTL,
		},
	})
	if err != nil {
		log.Fatalf("service launch failed: %v", err)
	}

	application, err := app.NewApp(app.Infrastructure{
		DB:          infra.DB(),
		Cache:       infra.Cache(),
		GoogleOAuth: infra.GoogleOAuth(),
		Jwt:         infra.Jwt(),
	})
	if err != nil {
		log.Fatalf("service launch failed: %v", err)
	}

	fiberApp := fiber.New(fiber.Config{
		StructValidator: validator.New(),
	})
	fiberApp.Use(middleware.ErrorHandler)
	fiberApp.Use(cors.New(c.CORS))

	docs.NewDocs().Register(fiberApp)
	rest.NewRestAPI(*application).Register(fiberApp)

	log.Fatal(fiberApp.Listen(":8080"))
}
