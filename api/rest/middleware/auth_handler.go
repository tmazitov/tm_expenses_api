package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

type UserService interface {
	Authenticate(token string) (int, error)
}

func AuthMiddleware(service UserService, isMock bool) fiber.Handler {

	if isMock {
		return func(ctx fiber.Ctx) error {

			ctx.Locals("userId", 1)

			return ctx.Next()
		}
	}

	return func(ctx fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := service.Authenticate(token)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		ctx.Locals("userId", userId)

		return ctx.Next()
	}
}
