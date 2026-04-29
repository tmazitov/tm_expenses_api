package expense

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
)

type expenseMonthlyCategoryStatRow struct {
	MonthNumber uint8           `bun:"month_number"`
	Total       decimal.Decimal `bun:"total"`
	Categories  json.RawMessage `bun:"categories"`
}

func (r *Repository) MonthlyCategoryStats(ctx context.Context, filters expense.ExpenseStatFilters) ([]*expense.ExpenseStat, error) {

	rows := []expenseMonthlyCategoryStatRow{}

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
		ColumnExpr(`(
			SELECT COALESCE(json_agg(json_build_object('id', sub.category_id, 'percent', sub.pct)), '[]'::json)
			FROM (
				SELECT
					category_id,
					(SUM(price) * 100 / NULLIF(SUM(SUM(price)) OVER (), 0))::int AS pct
				FROM expense
				WHERE DATE_TRUNC('month', created_at) = month_start
				  AND user_id = ?
				GROUP BY category_id
			) sub
		) AS categories`, filters.UserId()).
		Join("LEFT JOIN expense e ON DATE_TRUNC('month', e.created_at) = month_start AND e.user_id = ?", filters.UserId()).
		GroupExpr("month_start").
		OrderExpr("month_start ASC").
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

		stat, err := expense.NewExpenseStat(row.MonthNumber, row.Total, expense.MonthlyStat)
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
