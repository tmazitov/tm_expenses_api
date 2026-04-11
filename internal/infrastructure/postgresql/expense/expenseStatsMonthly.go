package expense

import (
	"context"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type expenseMonthlyStatRow struct {
	MonthNumber uint8           `bun:"month_number"`
	Total       decimal.Decimal `bun:"total"`
}

func (r *Repository) StatsMonthly(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {

	rows := []expenseMonthlyStatRow{}

	now := time.Now().UTC()
	currentMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	pageOffset := int(filters.Page()) * int(filters.Units())
	endPeriod := currentMonthStart.AddDate(0, -pageOffset, 0)
	startPeriod := endPeriod.AddDate(0, -int(filters.Units()-1), 0)

	err := r.db.NewSelect().
		TableExpr(`
			generate_series(
				DATE_TRUNC('month', ?::date),
				DATE_TRUNC('month', ?::date),
				'1 month'::interval
			) AS month_start
		`, startPeriod, endPeriod).
		ColumnExpr("EXTRACT(MONTH FROM month_start)::int AS month_number").
		ColumnExpr("COALESCE(SUM(e.price), 0) AS total").
		Join("LEFT JOIN expense e ON DATE_TRUNC('month', e.created_at) = month_start").
		GroupExpr("month_start").
		OrderExpr("month_start ASC").
		Scan(ctx, &rows)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	result := make([]*expense.ExpenseStat, 0, len(rows))
	for _, row := range rows {
		stat, err := expense.NewExpenseStat(row.MonthNumber, int(row.Total.IntPart()), expense.MonthlyStat)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}
		result = append(result, stat)
	}
	return result, nil
}
