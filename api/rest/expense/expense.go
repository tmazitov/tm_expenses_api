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

func (r *Router) Register(a *fiber.App) {
	r.group = a.Group("/expense")
}

func (r *Router) Use(middleware fiber.Handler) {
	r.group.Use(middleware)
}

func (r *Router) ApplyRoutes() {
	r.group.
		Post("/", r.Create()).
		Get("/", r.List()).
		Get("/stats", r.Stats())
}
