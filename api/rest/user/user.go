package user

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/ayda-order-service.git/internal/app/user"
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
		Post("/google", r.GoogleOAuth())
}
