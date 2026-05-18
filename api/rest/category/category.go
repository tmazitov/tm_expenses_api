package category

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/tm_expenses_api/internal/app/category"
)

type Router struct {
	service *category.Service
	group   fiber.Router
}

func NewRouter(service *category.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	r.group = a.Group("/category")
}

func (r *Router) Use(middleware fiber.Handler) {
	r.group.Use(middleware)
}

func (r *Router) ApplyRoutes() {
	r.group.Post("/", r.Create())
	r.group.Get("/", r.List())
}
