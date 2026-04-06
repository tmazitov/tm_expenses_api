package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx) error {
	err := ctx.Next()
	if err == nil {
		return nil
	}

	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	if code == fiber.StatusInternalServerError {
		slog.Error("internal server error",
			"method", ctx.Method(),
			"path", ctx.Path(),
			"err", err,
		)
	}

	return ctx.SendStatus(code)
}
