package expense

import "github.com/tmazitov/tm_expenses_api/internal/domain/expense"

type Service struct {
	repo expense.Repository
}

func NewService(r expense.Repository) *Service {
	return &Service{
		repo: r,
	}
}
