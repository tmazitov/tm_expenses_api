package expense

import (
	"context"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (er *Repository) GetById(ctx context.Context, id string) (*expense.Expense, error) {
	return nil, nil
}
