package user

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/tm_expenses_api/internal/app/user"
	userDomain "github.com/tmazitov/tm_expenses_api/internal/domain/user"
)

type GoogleOAuthRequest struct {
	IdToken string `json:"idToken" validate:"required,min=1"`
}

type GoogleOAuthResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// @Summary  Authenticate with Google OAuth
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    body body     GoogleOAuthRequest  true  "Google ID token"
// @Success  200  {object} GoogleOAuthResponse
// @Failure  400
// @Failure  500
// @Router   /auth/google [post]
func (r *Router) GoogleOAuth() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var req GoogleOAuthRequest

		if err := ctx.Bind().JSON(&req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		tokenPair, err := r.service.AuthWithGoogle(ctx, user.UserGoogleCredentials{
			IdToken: req.IdToken,
		})
		if err != nil {
			var userErr *userDomain.UserError

			if errors.As(err, &userErr) {
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(GoogleOAuthResponse{
			Access:  tokenPair.Access,
			Refresh: tokenPair.Refresh,
		})
	}
}
