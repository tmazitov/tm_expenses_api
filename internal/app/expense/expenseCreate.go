package expense

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type CreateExpenseForm struct {
	Name       string
	CategoryId string
	Date       time.Time
	Price      decimal.Decimal
}

type ExpenseOutput struct {
	Id         string
	Name       string
	CategoryId string
	Price      decimal.Decimal
	CreatedAt  time.Time
}

func (s *Service) Create(ctx context.Context, form CreateExpenseForm) (*ExpenseOutput, error) {

	e, err := expense.NewExpense(expense.ExpenseParams{
		Id:         uuid.NewString(),
		Name:       form.Name,
		CategoryId: form.CategoryId,
		Price:      form.Price,
		CreatedAt:  form.Date,
	})
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}

	return &ExpenseOutput{
		Id:         e.Id(),
		Name:       e.Name(),
		CategoryId: e.CategoryId(),
		Price:      e.Price(),
		CreatedAt:  e.CreatedAt(),
	}, nil
}
