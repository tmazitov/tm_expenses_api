package expense

import (
	"context"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type ListExpenseInput struct {
	Name  string
	Page  int
	Limit int
}

type ListExpenseOutput struct {
	Items []ExpenseListItem
}

type ExpenseListItem struct {
	Id   string
	Name string
}

func (s *Service) List(ctx context.Context, input ListExpenseInput) (*ListExpenseOutput, error) {

	filters, err := expense.NewListFilters(input.Name, input.Page, input.Limit)
	if err != nil {
		return nil, err
	}

	items, err := s.repo.List(ctx, *filters)
	if err != nil {
		return nil, err
	}

	result := &ListExpenseOutput{}
	for _, e := range items {
		result.Items = append(result.Items, ExpenseListItem{Id: e.Id(), Name: e.Name()})
	}
	return result, nil
}
