package user

import "context"

func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	return s.jwt.Dispose(ctx, refreshToken)
}
