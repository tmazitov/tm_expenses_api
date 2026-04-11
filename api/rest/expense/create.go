package expense

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
	expenseDomain "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type CreateExpenseRequest struct {
	Name       string          `json:"name" validate:"required,min=1,max=255"`
	CategoryId string          `json:"categoryId" validate:"omitempty,uuid"`
	Date       time.Time       `json:"date" validate:"required"`
	Price      decimal.Decimal `json:"price" validate:"decimal_min=0.01" swaggertype:"number"`
}

type CreateExpenseResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryId string    `json:"categoryId,omitempty"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"createdAt"`
}

// @Summary  Create expense
// @Tags     expense
// @Accept   json
// @Produce  json
// @Param    body body     CreateExpenseRequest  true  "Expense data"
// @Success  201  {object} CreateExpenseResponse
// @Failure  400
// @Failure  500
// @Router   /expense [post]
func (r *Router) Create() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var req CreateExpenseRequest

		if err := ctx.Bind().JSON(&req); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		output, err := r.service.Create(ctx, expense.CreateExpenseForm{
			Name:       req.Name,
			Price:      req.Price,
			CategoryId: req.CategoryId,
			Date:       req.Date,
		})
		if err != nil {

			var expenseErr *expenseDomain.ExpenseError

			if errors.As(err, &expenseErr) {
				log.Println("create expense error: ", err)
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusCreated).
			JSON(CreateExpenseResponse{
				Id:         output.Id,
				Name:       output.Name,
				CategoryId: output.CategoryId,
				Price:      output.Price.InexactFloat64(),
				CreatedAt:  output.CreatedAt,
			})
	}
}
