package expense

import (
	"context"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
)

type expenseWeeklyStatRow struct {
	WeekNumber uint8           `bun:"week_number"`
	Total      decimal.Decimal `bun:"total"`
}

func (r *Repository) StatsWeekly(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {

	rows := []expenseWeeklyStatRow{}

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	currentWeekStart := today.AddDate(0, 0, -int(today.Weekday()-time.Monday+7)%7)
	pageOffset := int(filters.Page()) * int(filters.Units()) * 7
	endPeriod := currentWeekStart.AddDate(0, 0, -pageOffset)
	startPeriod := endPeriod.AddDate(0, 0, -int(filters.Units()-1)*7)
	if endPeriod.Equal(today) {
		endPeriod = endPeriod.AddDate(0, 0, 1)
	}

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
		Join("LEFT JOIN expense e ON DATE_TRUNC('week', e.created_at) = week_start AND e.user_id = ?", filters.UserId()).
		GroupExpr("week_start").
		OrderExpr("week_start ASC").
		Scan(ctx, &rows)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	result := make([]*expense.ExpenseStat, 0, len(rows))
	for _, row := range rows {
		stat, err := expense.NewExpenseStat(row.WeekNumber, row.Total, expense.WeeklyStat)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}
		result = append(result, stat)
	}
	return result, nil
}
