package expense

import (
	"context"

	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
)

type ExpenseStatsInput struct {
	Variant string
	Units   int8
	Page    int
	UserId  int
}
type ExpenseStatsOutput struct {
	Items []*ExpenseStatsRecord
}

type ExpenseStatsRecord struct {
	Key        uint8
	Value      float64
	Categories []ExpenseStatsCategoryRecord
}

type ExpenseStatsCategoryRecord struct {
	Id      string
	Percent int
}

func (s *Service) Stats(ctx context.Context, input ExpenseStatsInput) (*ExpenseStatsOutput, error) {

	variant, err := expense.NewExpenseStatVariant(input.Variant)
	if err != nil {
		return nil, err
	}

	filters, err := expense.NewExpenseStatFilters(expense.ExpenseStatFiltersParams{
		Variant: variant,
		Units:   uint8(input.Units),
		Page:    input.Page,
		UserId:  input.UserId,
	})
	if err != nil {
		return nil, err
	}

	output := ExpenseStatsOutput{
		Items: []*ExpenseStatsRecord{},
	}

	var stats []*expense.ExpenseStat

	if filters.Variant() == expense.WeeklyStat {
		stats, err = s.repo.WeeklyStats(ctx, *filters)
	} else if filters.Variant() == expense.WeeklyCategoryStat {
		stats, err = s.repo.WeeklyCategoryStats(ctx, *filters)
	} else if filters.Variant() == expense.MonthlyStat {
		stats, err = s.repo.MonthlyStats(ctx, *filters)
	} else if filters.Variant() == expense.MonthlyCategoryStat {
		stats, err = s.repo.MonthlyCategoryStats(ctx, *filters)
	}
	if err != nil {
		return nil, err
	}
	for _, stat := range stats {

		categories := stat.Categories()
		var categoryRecords []ExpenseStatsCategoryRecord

		if len(categories) != 0 {
			categoryRecords = make([]ExpenseStatsCategoryRecord, 0, len(categories))
			for _, categoryStat := range categories {
				categoryRecords = append(categoryRecords, ExpenseStatsCategoryRecord{
					Id:      categoryStat.Id(),
					Percent: categoryStat.Percent(),
				})
			}
		}

		output.Items = append(output.Items, &ExpenseStatsRecord{
			Key:        stat.Key(),
			Value:      stat.Value().InexactFloat64(),
			Categories: categoryRecords,
		})
	}

	return &output, nil
}
