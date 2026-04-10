package category

import "context"

type Repository interface {
	Create(ctx context.Context, c *Category) error
	List(ctx context.Context) ([]*Category, error)
}
