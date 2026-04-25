package user

import "context"

func (s *Service) Refresh(ctx context.Context, token string) (*JwtTokenPair, error) {

	access, refresh, err := s.jwt.Refresh(ctx, token)
	if err != nil {
		return nil, err
	}

	return &JwtTokenPair{
		Access:  access,
		Refresh: refresh,
	}, nil
}
