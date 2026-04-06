package expense

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/tmazitov/ayda-order-service.git/internal/app/expense"
	expenseDomain "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type CreateExpenseRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type CreateExpenseResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// @Summary  Create expense
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
			Name: req.Name,
		})
		if err != nil {
			var expenseErr *expenseDomain.ExpenseError

			if errors.As(err, &expenseErr) {
				return ctx.SendStatus(fiber.StatusBadRequest)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusCreated).
			JSON(CreateExpenseResponse{
				Id:   output.Id,
				Name: output.Name,
			})
	}
}
