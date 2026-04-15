package expense

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type Mock struct{}

func (m Mock) Create(ctx context.Context, p *expense.Expense) error             { return nil }
func (m Mock) GetById(ctx context.Context, id string) (*expense.Expense, error) { return nil, nil }
func (m Mock) List(ctx context.Context, filters expense.ListFilters) ([]*expense.Expense, error) {

	instance, _ := expense.NewExpense(expense.ExpenseParams{
		Id:         uuid.NewString(),
		Name:       "expense",
		Price:      decimal.New(1, 10),
		CreatedAt:  time.Now(),
		CategoryId: uuid.NewString(),
	})

	return []*expense.Expense{
		instance,
	}, nil
}
func (m Mock) StatsWeekly(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}
func (m Mock) StatsMonthly(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {
	return []*expense.ExpenseStat{}, nil
}
