package expense

import "context"

type Repository interface {
	Create(ctx context.Context, p *Expense) error
	List(ctx context.Context, filters ListFilters) ([]*Expense, error)
	StatsWeekly(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
	StatsMonthly(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
}
