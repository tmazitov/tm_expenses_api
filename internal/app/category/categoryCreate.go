package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

type CategoryCreateForm struct {
	Name  string
	Icon  string
	Color string
}

type CategoryCreateOutput struct {
	Id    string
	Name  string
	Icon  string
	Color string
}

func (s *Service) Create(ctx context.Context, form CategoryCreateForm) (*CategoryCreateOutput, error) {

	params := category.CategoryParams{
		Id:   uuid.NewString(),
		Name: form.Name,
		Icon: form.Icon,
	}

	if len(form.Color) != 0 {
		color, err := category.NewColor(form.Color)
		if err != nil {
			return nil, err
		}
		params.Color = &color
	}

	c, err := category.NewCategory(params)
	if err != nil {
		return nil, err
	}

	err = s.repository.Create(ctx, c)
	if err != nil {
		return nil, err
	}

	output := &CategoryCreateOutput{
		Id:   c.Id(),
		Name: c.Name(),
		Icon: c.Icon(),
	}

	if c.Color() != nil {
		output.Color = c.Color().Hex()
	}

	return output, nil
}
