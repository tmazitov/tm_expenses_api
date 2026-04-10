package expense

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
	expenseDomain "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type ListExpenseQuery struct {
	// Date *time.Time `query:"date"`
	Name     string `query:"name"`
	Page     int    `query:"page" validate:"min=0"`
	Limit    int    `query:"limit" validate:"min=0,max=100"`
	Date     string `query:"date" validate:"required,date=02.01.2006"`
	CategoryId string `query:"category" validate:"omitempty,uuid"`
}

type ListExpenseResponse struct {
	Items []ListExpenseItem `json:"items"`
}

type ListExpenseItem struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryId string    `json:"categoryId,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	Price      float64   `json:"price"`
}

// @Summary  List expenses
// @Tags     expense
// @Produce  json
// @Param    name     query    string  false  "Filter by name"
// @Param    date     query    string  false  "Filter by date (dd.mm.yyyy)"
// @Param    page     query    int     false  "Page number (0-based)"  minimum(0)
// @Param    limit    query    int     false  "Items per page"         minimum(1) maximum(100)
// @Param    category query    string  false  "Filter by category UUID"
// @Success  200      {object} ListExpenseResponse
// @Failure  400
// @Failure  500
// @Router   /expense [get]
func (r *Router) List() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var filters ListExpenseQuery
		if err := ctx.Bind().Query(&filters); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		date, err := time.Parse("02.01.2006", filters.Date)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		list, err := r.service.List(ctx, expense.ListExpenseInput{
			Name:       filters.Name,
			Limit:      filters.Limit,
			Page:       filters.Page,
			CategoryId: filters.CategoryId,
			Date:       date,
		})
		if err != nil {
			var expenseErr *expenseDomain.ExpenseError
			log.Println(err)
			if errors.As(err, &expenseErr) {
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		var result = &ListExpenseResponse{
			Items: make([]ListExpenseItem, 0, len(list.Items)),
		}

		for _, item := range list.Items {
			result.Items = append(result.Items, ListExpenseItem{
				Id:         item.Id,
				Name:       item.Name,
				CreatedAt:  item.CreatedAt,
				Price:      item.Price.InexactFloat64(),
				CategoryId: item.CategoryId,
			})
		}

		return ctx.JSON(result)
	}
}
