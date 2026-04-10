package category

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/internal/app/category"
	domain "github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

type CreateCategoryRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Color string `json:"color" validate:"omitempty,hexcolor"`
	Icon  string `json:"icon" validate:"max=100"`
}

type CreateCategoryResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

// @Summary  Create category
// @Tags     category
// @Accept   json
// @Produce  json
// @Param    body body     CreateCategoryRequest  true  "Category data"
// @Success  201  {object} CreateCategoryResponse
// @Failure  400
// @Failure  500
// @Router   /category [post]
func (r *Router) Create() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var req CreateCategoryRequest

		if err := ctx.Bind().JSON(&req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		c, err := r.service.Create(ctx, category.CategoryCreateForm{
			Name:  req.Name,
			Color: req.Color,
			Icon:  req.Icon,
		})
		if err != nil {

			var expenseErr *domain.CategoryError

			if errors.As(err, &expenseErr) {
				log.Println("create category error: ", err)
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusCreated).
			JSON(CreateCategoryResponse{
				Id:    c.Id,
				Name:  c.Name,
				Color: c.Color,
				Icon:  c.Icon,
			})
	}
}
