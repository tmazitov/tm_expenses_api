package category

import (
	"context"

	"github.com/gofiber/fiber/v3"

	app "github.com/tmazitov/ayda-order-service.git/internal/app/category"
)

type CategoryService interface {
	Create(ctx context.Context, input app.CategoryCreateForm) (*app.CategoryCreateOutput, error)
	List(ctx context.Context) ([]app.CategoryListItem, error)
}

type Router struct {
	service CategoryService
}

func NewRouter(service CategoryService) *Router {
	return &Router{
		service: service,
	}
}

func (r *Router) Register(a *fiber.App) {
	a.Group("/category").
		Post("/", r.Create()).
		Get("/", r.List())
}
