package expense

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type expenseWeeklyStatRow struct {
	WeekNumber uint8           `bun:"week_number"`
	Total      decimal.Decimal `bun:"total"`
}

func (r *Repository) StatsWeekly(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {

	rows := []expenseWeeklyStatRow{}

	now := time.Now()
	currentWeekStart := now.AddDate(0, 0, -int(now.Weekday()-time.Monday+7)%7)
	pageOffset := int(filters.Page()) * int(filters.Units()) * 7
	endPeriod := currentWeekStart.AddDate(0, 0, -pageOffset)
	startPeriod := endPeriod.AddDate(0, 0, -int(filters.Units()-1)*7)

	err := r.db.NewSelect().
		TableExpr(`
			generate_series(
				DATE_TRUNC('week', ?::date),
				DATE_TRUNC('week', ?::date),
				'1 week'::interval
			) AS week_start
		`, startPeriod, endPeriod).
		ColumnExpr("EXTRACT(WEEK FROM week_start)::int AS week_number").
		ColumnExpr("COALESCE(SUM(e.price), 0) AS total").
		Join("LEFT JOIN expense e ON DATE_TRUNC('week', e.created_at) = week_start").
		GroupExpr("week_start").
		OrderExpr("week_start ASC").
		Scan(ctx, &rows)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}
	fmt.Printf("units: %d\n", filters.Units())
	fmt.Printf("start: %s\nend: %s\nrecords: %+v\n", startPeriod, endPeriod, rows)

	result := make([]*expense.ExpenseStat, 0, len(rows))
	for _, row := range rows {
		stat, err := expense.NewExpenseStat(row.WeekNumber, int(row.Total.IntPart()), expense.WeeklyStat)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}
		result = append(result, stat)
	}
	return result, nil
}
