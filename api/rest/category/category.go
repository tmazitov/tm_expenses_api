package category

import (
	"github.com/gofiber/fiber/v3"

	"github.com/tmazitov/ayda-order-service.git/internal/app/category"
)

type Router struct {
	service *category.Service
}

func NewRouter(service *category.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	a.Group("/category").
		Post("/", r.Create()).
		Get("/", r.List())
}
