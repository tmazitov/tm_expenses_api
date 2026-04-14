package category

import (
	"context"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/category"
)

type Mock struct{}

func (m Mock) Create(ctx context.Context, c *category.Category) error { return nil }
func (m Mock) List(ctx context.Context) ([]*category.Category, error) { return nil, nil }
