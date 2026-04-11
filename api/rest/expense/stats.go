package expense

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
	expenseDomain "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type ExpenseStatsQuery struct {
	Variant string `query:"variant" validate:"required"`
	Units   int8   `query:"units" validate:"min=1,max=6"`
	Page    int    `query:"page" validate:"min=0"`
}

type ExpenseStatItem struct {
	Key   uint8 `json:"key"`
	Value int   `json:"value"`
}

type ExpenseStatsResponse struct {
	Items []*ExpenseStatItem `json:"items"`
}

// @Summary  Get expense stats
// @Tags     expense
// @Produce  json
// @Param    variant  query    string  true   "Stat variant (weekly, monthly)"
// @Param    units    query    int     false  "Number of periods to look back (1–6)"  minimum(1) maximum(6)
// @Param    page     query    int     false  "Page number (0-based)"                 minimum(0)
// @Success  200      {object} ExpenseStatsResponse
// @Failure  400
// @Failure  500
// @Router   /expense/stats [get]
func (r *Router) Stats() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var filters ExpenseStatsQuery
		if err := ctx.Bind().Query(&filters); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		stats, err := r.service.Stats(ctx, expense.ExpenseStatsInput{
			Variant: filters.Variant,
			Units:   filters.Units,
			Page:    filters.Page,
		})
		if err != nil {
			var expenseErr *expenseDomain.ExpenseError
			log.Println(err)
			if errors.As(err, &expenseErr) {
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		response := ExpenseStatsResponse{
			Items: make([]*ExpenseStatItem, 0, len(stats.Items)),
		}

		for _, item := range stats.Items {
			response.Items = append(response.Items, &ExpenseStatItem{
				Key:   uint8(item.Key),
				Value: item.Value,
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
