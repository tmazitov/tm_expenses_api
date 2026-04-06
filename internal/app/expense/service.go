package expense

import "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"

type Service struct {
	repo expense.Repository
}

func NewService(r expense.Repository) *Service {
	return &Service{
		repo: r,
	}
}
