package expense

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, p *Expense) error
	List(ctx context.Context, filters ListFilters) ([]*Expense, error)
	WeeklyStats(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
	WeeklyCategoryStats(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
	MonthlyStats(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
	MonthlyCategoryStats(ctx context.Context, filters ExpenseStatFilters) ([]*ExpenseStat, error)
}
