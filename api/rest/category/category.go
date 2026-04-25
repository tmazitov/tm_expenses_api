package category

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/ayda-order-service.git/internal/app/category"
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

func (r *Router) Register(a *fiber.App) *Router {
	r.group = a.Group("/category")
	return r
}

func (r *Router) Use(middleware fiber.Handler) *Router {
	r.group.Use(middleware)
	return r
}

func (r *Router) Routes() *Router {
	r.group.Post("/", r.Create())
	r.group.Get("/", r.List())
	return r
}
