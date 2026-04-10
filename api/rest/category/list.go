package category

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	domain "github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

type CategoryListResponse struct {
	Items []CategoryListItem `json:"items"`
}

type CategoryListItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

// @Summary  List categories
// @Tags     category
// @Produce  json
// @Success  200  {object} CategoryListResponse
// @Failure  400
// @Failure  500
// @Router   /category [get]
func (r *Router) List() fiber.Handler {
	return func(ctx fiber.Ctx) error {

		categories, err := r.service.List(ctx)
		if err != nil {

			var expenseErr *domain.CategoryError

			if errors.As(err, &expenseErr) {
				log.Println("create category error: ", err)
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		output := CategoryListResponse{
			Items: make([]CategoryListItem, 0, len(categories)),
		}

		for _, c := range categories {
			output.Items = append(output.Items, CategoryListItem{
				Id:    c.Id,
				Name:  c.Name,
				Color: c.Color,
				Icon:  c.Icon,
			})
		}

		return ctx.Status(200).JSON(output)
	}
}
