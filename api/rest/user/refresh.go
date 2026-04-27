package user

import "github.com/gofiber/fiber/v3"

type RefreshRequest struct {
	RefreshToken string `json:"refresh" validate:"required,min=1"`
}

type RefreshResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// @Summary  Refresh access token
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    body body     RefreshRequest  true  "Refresh token"
// @Success  200  {object} RefreshResponse
// @Failure  401
// @Router   /auth/refresh [post]
func (r *Router) Refresh() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		var req RefreshRequest

		if err := ctx.Bind().JSON(&req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		tokenPair, err := r.service.Refresh(ctx, req.RefreshToken)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		return ctx.Status(fiber.StatusOK).JSON(RefreshResponse{
			Access:  tokenPair.Access,
			Refresh: tokenPair.Refresh,
		})
	}
}
