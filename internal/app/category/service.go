package category

import "github.com/tmazitov/ayda-order-service.git/internal/domain/category"

type Service struct {
	repository category.Repository
}

func NewService(repository category.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
