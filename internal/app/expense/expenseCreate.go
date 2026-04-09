package expense

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type CreateExpenseForm struct {
	Name  string
	Price decimal.Decimal
}

type ExpenseOutput struct {
	Id        string
	Name      string
	Price     decimal.Decimal
	CreatedAt time.Time
}

func (s *Service) Create(ctx context.Context, input CreateExpenseForm) (*ExpenseOutput, error) {
	e, err := expense.NewExpense(uuid.NewString(), input.Name, input.Price)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}

	return &ExpenseOutput{
		Id:        e.Id(),
		Name:      e.Name(),
		Price:     e.Price(),
		CreatedAt: e.CreatedAt(),
	}, nil
}
