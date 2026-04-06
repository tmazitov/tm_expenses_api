package expense

import (
	"context"

	"github.com/google/uuid"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type CreateExpenseForm struct {
	Name string
}

type ExpenseOutput struct {
	Id   string
	Name string
}

func (s *Service) Create(ctx context.Context, input CreateExpenseForm) (*ExpenseOutput, error) {
	e, err := expense.NewExpense(uuid.New(), input.Name)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}

	return &ExpenseOutput{
		Id:   e.Id(),
		Name: e.Name(),
	}, nil
}
