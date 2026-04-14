package expense

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
)

type Router struct {
	service *expense.Service
}

func NewRouter(service *expense.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	a.Group("/expense").
		Post("/", r.Create()).
		Get("/", r.List()).
		Get("/stats", r.Stats())
}
