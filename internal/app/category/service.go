package category

import "github.com/tmazitov/tm_expenses_api/internal/domain/category"

type Service struct {
	repository category.Repository
}

func NewService(repository category.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
