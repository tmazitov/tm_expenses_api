package expense

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/tm_expenses_api/internal/app/expense"
)

type Router struct {
	service *expense.Service
	group   fiber.Router
}

func NewRouter(service *expense.Service) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) *Router {
	r.group = a.Group("/expense")
	return r
}

func (r *Router) Use(middleware fiber.Handler) *Router {
	r.group.Use(middleware)
	return r
}

func (r *Router) Routes() *Router {
	r.group.
		Post("/", r.Create()).
		Get("/", r.List()).
		Get("/stats", r.Stats())
	return r
}
