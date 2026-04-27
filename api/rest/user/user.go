package user

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/tm_expenses_api/internal/app/user"
)

type Router struct {
	service   *user.Service
	authGroup fiber.Router
	userGroup fiber.Router
}

func NewRouter(service *user.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) *Router {
	r.authGroup = a.Group("/auth")
	r.userGroup = a.Group("/user")
	return r
}

func (r *Router) Use(middleware fiber.Handler) *Router {
	r.userGroup.Use(middleware)
	return r
}

func (r *Router) Routes() *Router {
	r.authGroup.
		Post("/google", r.GoogleOAuth()).
		Post("/refresh", r.Refresh())

	r.userGroup.
		Get("", r.Profile())

	return r
}
