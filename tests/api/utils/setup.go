package utils

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/api/rest"
	"github.com/tmazitov/ayda-order-service.git/internal/app"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
	"github.com/tmazitov/ayda-order-service.git/pkg/validator"
)

func SetupAppInstance() *fiber.App {
	svc := app.NewApp(postgresql.Mock{})

	app := fiber.New(fiber.Config{
		StructValidator: validator.New(),
	})
	rest.NewRestAPI(*svc).Register(app)

	return app
}
