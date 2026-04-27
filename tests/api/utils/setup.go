package utils

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/tm_expenses_api/api/rest"
	"github.com/tmazitov/tm_expenses_api/internal/app"
	"github.com/tmazitov/tm_expenses_api/internal/infrastructure/postgresql"
	"github.com/tmazitov/tm_expenses_api/pkg/validator"
)

func SetupAppInstance() *fiber.App {
	svc := app.NewApp(postgresql.Mock{})

	app := fiber.New(fiber.Config{
		StructValidator: validator.New(),
	})
	rest.NewRestAPI(*svc).Register(app)

	return app
}
