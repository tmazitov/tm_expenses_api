package mock

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/tm_expenses_api/api/rest"
	"github.com/tmazitov/tm_expenses_api/internal/app"
	"github.com/tmazitov/tm_expenses_api/pkg/validator"
)

func NewMockApp() *fiber.App {
	svc, _ := app.NewApp(app.Infrastructure{
		DB:          mockDB{},
		Cache:       mockCache{},
		Jwt:         mockJwtProvider{},
		GoogleOAuth: mockGoogleOAuthProvider{},
	})

	app := fiber.New(fiber.Config{
		StructValidator: validator.New(),
	})
	rest.NewRestAPI(*svc, rest.UnsafeMode).Register(app)

	return app
}
