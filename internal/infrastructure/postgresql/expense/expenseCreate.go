package expense

import (
	"context"
	"errors"

	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
)

func (r *Repository) Create(ctx context.Context, e *expense.Expense) error {

	model := newExpenseModel(e)

	_, err := r.db.NewInsert().
		Model(model).
		Exec(ctx)

	if err != nil {
		return errors.Join(ErrInsertionFailed, err)
	}

	return nil
}
