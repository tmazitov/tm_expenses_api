package expense

import (
	"context"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

type ExpenseStatsInput struct {
	Variant string
	Units   int8
	Page    int
}
type ExpenseStatsOutput struct {
	Items []*ExpenseStatsRecord
}

type ExpenseStatsRecord struct {
	Key   uint8
	Value int
}

func (s *Service) Stats(ctx context.Context, input ExpenseStatsInput) (*ExpenseStatsOutput, error) {

	variant, err := expense.NewExpenseStatVariant(input.Variant)
	if err != nil {
		return nil, err
	}

	params := expense.ExpenseStatFiltersParams{
		Variant: variant,
		Units:   uint8(input.Units),
		Page:    input.Page,
	}

	filters, err := expense.NewExpenseStatFilters(params)
	if err != nil {
		return nil, err
	}

	output := ExpenseStatsOutput{
		Items: []*ExpenseStatsRecord{},
	}

	if filters.Variant() == expense.WeeklyStat {
		stats, err := s.repo.StatsWeekly(ctx, *filters)
		if err != nil {
			return nil, err
		}

		for _, stat := range stats {
			output.Items = append(output.Items, &ExpenseStatsRecord{
				Key:   stat.Key(),
				Value: stat.Value(),
			})
		}
	}

	return &output, nil
}
