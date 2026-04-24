package user

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

type UserGoogleCredentials struct {
	IdToken string
}

type UserCreateOutput struct {
	Access  string
	Refresh string
}

func (s *Service) AuthWithGoogle(ctx context.Context, credentials UserGoogleCredentials) (*UserCreateOutput, error) {

	claims, err := s.googleOauthProvider.ValidateToken(ctx, credentials.IdToken)
	if err != nil {
		return nil, errors.Join(ErrInvalidCredentials, err)
	}

	params, err := claimsToParams(claims)
	if err != nil {
		return nil, errors.Join(ErrGoogleOauthFailed, err)
	}

	u, err := s.repo.GetBySub(ctx, user.GoogleOauth, params.Sub)
	if err != nil {
		return nil, err
	}

	// If user is undefined, create a new one
	if u == nil {
		params.AuthMethod = user.GoogleOauth
		u, err = user.NewUser(params)
		if err != nil {
			return nil, err
		}

		if u, err = s.repo.Create(ctx, u); err != nil {
			return nil, err
		}
	} else {
		u, err = u.WithUpdatedInfo(params)
		if err != nil {
			return nil, err
		}

		if err = s.repo.Update(ctx, u); err != nil {
			return nil, err
		}
	}

	// Create access and refresh tokens
	access, refresh, err := s.jwt.CreateTokenPair(ctx, newUserClaims(u))
	if err != nil {
		return nil, err
	}

	return &UserCreateOutput{
		Access:  access,
		Refresh: refresh,
	}, nil
}
