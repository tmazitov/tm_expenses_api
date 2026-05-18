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

func (r *Router) Register(a *fiber.App) {
	r.authGroup = a.Group("/auth")
	r.userGroup = a.Group("/user")
}

func (r *Router) Use(middleware fiber.Handler) {
	r.userGroup.Use(middleware)
}

func (r *Router) ApplyRoutes() {
	r.authGroup.
		Post("/google", r.GoogleOAuth()).
		Post("/refresh", r.Refresh())

	r.userGroup.
		Get("", r.Profile()).
		Post("/logout", r.Logout())
}
