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
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
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

	db, err := postgresql.NewDatabase("./db/migrations", c.DB)
	if err != nil {
		log.Fatalf("service launch failed: %v", err)
	}
	defer db.Close()

	application := app.NewApp(db)

	fiberApp := fiber.New(fiber.Config{
		StructValidator: validator.New(),
	})
	fiberApp.Use(middleware.ErrorHandler)
	fiberApp.Use(cors.New(c.CORS))

	docs.NewDocs().Register(fiberApp)
	rest.NewRestAPI(*application).Register(fiberApp)

	log.Fatal(fiberApp.Listen(":8080"))
}
