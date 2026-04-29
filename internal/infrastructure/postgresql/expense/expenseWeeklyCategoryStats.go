package expense

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
)

type categoryExpenseWeeklyStatRow struct {
	WeekNumber uint8           `bun:"week_number"`
	Total      decimal.Decimal `bun:"total"`
	Categories json.RawMessage `bun:"categories"`
}

type categoryItemRow struct {
	Id      string `bun:"id"`
	Percent int    `bun:"percent"`
}

func (r *Repository) WeeklyCategoryStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {

	rows := []categoryExpenseWeeklyStatRow{}

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
		ColumnExpr(`(
			SELECT COALESCE(json_agg(json_build_object('id', sub.category_id, 'percent', sub.pct)), '[]'::json)
			FROM (
				SELECT
					category_id,
					(SUM(price) * 100 / NULLIF(SUM(SUM(price)) OVER (), 0))::int AS pct
				FROM expense
				WHERE DATE_TRUNC('week', created_at) = week_start
				  AND user_id = ?
				GROUP BY category_id
			) sub
		) AS categories`, filters.UserId()).
		Join("LEFT JOIN expense e ON DATE_TRUNC('week', e.created_at) = week_start AND e.user_id = ?", filters.UserId()).
		GroupExpr("week_start").
		OrderExpr("week_start ASC").
		Scan(ctx, &rows)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	result := make([]*expense.ExpenseStat, 0, len(rows))
	for _, row := range rows {

		categoryRows := []categoryItemRow{}

		err := json.Unmarshal(row.Categories, &categoryRows)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}

		fmt.Printf("cats: %+v\n", categoryRows)

		var categories []expense.ExpenseStatCategory

		if len(categoryRows) != 0 {
			categories = make([]expense.ExpenseStatCategory, 0, len(categoryRows))
			for _, catRow := range categoryRows {

				category, err := expense.NewExpenseStatCategory(catRow.Id, catRow.Percent)
				if err != nil {
					return nil, errors.Join(ErrSelectionFailed, err)
				}

				categories = append(categories, category)
			}
		}

		stat, err := expense.NewExpenseStat(row.WeekNumber, row.Total, expense.WeeklyStat)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}

		if len(categories) != 0 {
			stat = stat.WithCategories(categories)
		}

		result = append(result, stat)
	}
	return result, nil
}
