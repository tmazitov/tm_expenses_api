package user

import (
	"context"
	"fmt"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

type GoogleOAuthProvider interface {
	ValidateToken(ctx context.Context, idToken string) (map[string]any, error)
}

type Service struct {
	googleOauthProvider GoogleOAuthProvider
	repo                user.Repository
	jwt                 JwtProvider
}

type ServiceParams struct {
	GoogleOAuthProvider GoogleOAuthProvider
	Repo                user.Repository
	Jwt                 JwtProvider
}

func (p ServiceParams) validate() error {

	if p.GoogleOAuthProvider == nil {
		return fmt.Errorf("%w : google oauth provider is nil.", ErrInvalidParams)
	}

	if p.Repo == nil {
		return fmt.Errorf("%w : repo is nil.", ErrInvalidParams)
	}

	if p.Jwt == nil {
		return fmt.Errorf("%w : repo is nil.", ErrInvalidParams)
	}

	return nil
}

func NewService(params ServiceParams) (*Service, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &Service{
		googleOauthProvider: params.GoogleOAuthProvider,
		repo:                params.Repo,
		jwt:                 params.Jwt,
	}, nil
}

type JwtTokenPair struct {
	Access  string
	Refresh string
}
