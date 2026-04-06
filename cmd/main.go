package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	restExpense "github.com/tmazitov/ayda-order-service.git/api/rest/expense"
	"github.com/tmazitov/ayda-order-service.git/api/rest/middleware"
	"github.com/tmazitov/ayda-order-service.git/config"
	"github.com/tmazitov/ayda-order-service.git/internal/app"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
)

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

	fiberApp := fiber.New()
	fiberApp.Use(middleware.ErrorHandler)

	restExpense.NewRouter(application.ExpenseService()).Register(fiberApp)

	log.Fatal(fiberApp.Listen(":8080"))
}
