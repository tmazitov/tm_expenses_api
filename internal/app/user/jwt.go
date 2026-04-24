package user

import "context"

type JwtProvider interface {
	CreateTokenPair(ctx context.Context, claims map[string]any) (string, string, error)
	Refresh(ctx context.Context, token string) (string, string, error)
	VerifyAccess(token string) (map[string]any, error)
	VerifyRefresh(ctx context.Context, token string) (map[string]any, error)
}
