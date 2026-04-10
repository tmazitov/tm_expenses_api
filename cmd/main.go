package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/tmazitov/ayda-order-service.git/api/docs"
	restCategory "github.com/tmazitov/ayda-order-service.git/api/rest/category"
	restExpense "github.com/tmazitov/ayda-order-service.git/api/rest/expense"
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

	docs.NewRouter().Register(fiberApp)
	restExpense.NewRouter(application.ExpenseService()).Register(fiberApp)
	restCategory.NewRouter(application.CategoryService()).Register(fiberApp)

	log.Fatal(fiberApp.Listen(":8080"))
}
