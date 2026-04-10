package category

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

func (r *Repository) List(ctx context.Context) ([]*category.Category, error) {

	var models = []*CategoryModel{}

	err := r.db.NewSelect().
		Model(&models).
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	var categories = make([]*category.Category, 0, len(models))
	for _, model := range models {

		params := category.CategoryParams{
			Id:   model.Id,
			Name: model.Name,
		}

		if model.Color != nil {
			color, err := category.RestoreColor(*model.Color)
			if err != nil {
				return nil, errors.Join(ErrSelectionFailed, err)
			}
			params.Color = &color
		}

		if model.Icon != nil {
			params.Icon = *model.Icon
		}

		c, err := category.NewCategory(params)
		if err != nil {
			return nil, errors.Join(ErrSelectionFailed, err)
		}

		categories = append(categories, c)
	}

	return categories, nil
}
