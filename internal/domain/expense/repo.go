package expense

import "context"

type Repository interface {
	Create(ctx context.Context, p *Expense) error
	GetById(ctx context.Context, id string) (*Expense, error)
	List(ctx context.Context, filters ListFilters) ([]*Expense, error)
	StatsWeekly(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
}
