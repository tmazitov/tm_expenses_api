package category

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

func (r *Repository) Create(ctx context.Context, c *category.Category) error {

	model := NewCategoryModel(c)

	_, err := r.db.NewInsert().
		Model(model).
		Exec(ctx)

	if err != nil {
		return errors.Join(ErrInsertionFailed, nil)
	}

	return nil
}
