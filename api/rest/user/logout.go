package user

import "github.com/gofiber/fiber/v3"

type LogoutRequest struct {
	Refresh string `json:"refresh"`
}

func (r *Router) Logout() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		var req LogoutRequest

		if err := ctx.Bind().JSON(&req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		err := r.service.Logout(ctx, req.Refresh)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
