package expense

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type ListExpenseInput struct {
	Name  string
	Page  int
	Limit int
	Date  time.Time
}

type ListExpenseOutput struct {
	Items []ExpenseListItem
}

type ExpenseListItem struct {
	Id        string
	Name      string
	Price     decimal.Decimal
	CreatedAt time.Time
}

func NewExpenseListItem(e *expense.Expense) ExpenseListItem {
	return ExpenseListItem{
		Id:        e.Id(),
		Name:      e.Name(),
		Price:     e.Price(),
		CreatedAt: e.CreatedAt(),
	}
}

func (s *Service) List(ctx context.Context, input ListExpenseInput) (*ListExpenseOutput, error) {

	filters, err := expense.NewListFilters(expense.ListFiltersParams{
		Name:  input.Name,
		Date:  input.Date,
		Page:  input.Page,
		Limit: input.Limit,
	})
	if err != nil {
		return nil, err
	}

	items, err := s.repo.List(ctx, *filters)
	if err != nil {
		return nil, err
	}

	result := &ListExpenseOutput{}
	for _, e := range items {
		result.Items = append(result.Items, NewExpenseListItem(e))
	}
	return result, nil
}
