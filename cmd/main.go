package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/tmazitov/tm_expenses_api/api/docs"
	"github.com/tmazitov/tm_expenses_api/api/rest"
	"github.com/tmazitov/tm_expenses_api/api/rest/middleware"
	"github.com/tmazitov/tm_expenses_api/config"
	"github.com/tmazitov/tm_expenses_api/internal/app"
	infra "github.com/tmazitov/tm_expenses_api/internal/infrastructure"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/google"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/jwt"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/redis"
	"github.com/tmazitov/tm_expenses_api/pkg/validator"
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
			Addr:     c.Cache.Addr,
			DB:       c.Cache.DB,
			Password: c.Cache.Password,
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
	rest.NewRestAPI(*application, rest.SecuredMode).Register(fiberApp)

	log.Fatal(fiberApp.Listen(":8080"))
}
