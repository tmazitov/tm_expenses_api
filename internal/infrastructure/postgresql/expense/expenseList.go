package expense

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
)

func (r *Repository) List(ctx context.Context, filters expense.ListFilters) ([]*expense.Expense, error) {
	var models []expenseModel

	q := r.db.NewSelect().
		Model(&models).
		Limit(filters.Limit()).Offset(filters.Page() * filters.Limit())

	if len(filters.Name()) != 0 {
		q = q.Where("name ILIKE ?", "%"+filters.Name()+"%")
	}

	if err := q.Scan(ctx); err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	result := make([]*expense.Expense, 0, len(models))
	for _, m := range models {
		e, err := expense.RestoreExpense(m.Id, m.Name, m.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}
