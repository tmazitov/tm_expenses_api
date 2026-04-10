package expense

import (
	"context"

	"github.com/gofiber/fiber/v3"
	app "github.com/tmazitov/ayda-order-service.git/internal/app/expense"
)

type ExpenseService interface {
	Create(ctx context.Context, input app.CreateExpenseForm) (*app.ExpenseOutput, error)
	List(ctx context.Context, input app.ListExpenseInput) (*app.ListExpenseOutput, error)
}

type Router struct {
	service ExpenseService
}

func NewRouter(service ExpenseService) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	a.Group("/expense").
		Post("/", r.Create()).
		Get("/", r.List())
}
