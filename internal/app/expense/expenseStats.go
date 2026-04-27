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
	Key   uint8
	Value float64
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
		stats, err = s.repo.StatsWeekly(ctx, *filters)
		if err != nil {
			return nil, err
		}
	} else if filters.Variant() == expense.MonthlyStat {
		stats, err = s.repo.StatsMonthly(ctx, *filters)
		if err != nil {
			return nil, err
		}
	}
	for _, stat := range stats {
		output.Items = append(output.Items, &ExpenseStatsRecord{
			Key:   stat.Key(),
			Value: stat.Value().InexactFloat64(),
		})
	}

	return &output, nil
}
