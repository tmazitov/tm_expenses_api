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
		Where("created_at::date = ?", filters.Date()).
		Limit(filters.Limit()).Offset(filters.Page() * filters.Limit())

	if len(filters.Name()) != 0 {
		q = q.Where("name ILIKE ?", "%"+filters.Name()+"%")
	}

	if len(filters.CategoryId()) != 0 {
		q = q.Where("category_id = ?", filters.CategoryId())
	}

	if err := q.Scan(ctx); err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	result := make([]*expense.Expense, 0, len(models))
	for _, m := range models {
		e, err := expense.NewExpense(expense.ExpenseParams{
			Id:         m.Id,
			Name:       m.Name,
			Price:      m.Price,
			CreatedAt:  m.CreatedAt,
			CategoryId: m.CategoryId,
		})
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}
