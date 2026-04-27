package user

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/tm_expenses_api/internal/app/user"
)

type Router struct {
	service *user.Service
}

func NewRouter(service *user.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	a.Group("/auth").
		Post("/google", r.GoogleOAuth()).
		Post("/refresh", r.Refresh())
}
