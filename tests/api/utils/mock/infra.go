package mock

import (
	"context"
	"time"

	"github.com/tmazitov/tm_expenses_api/internal/domain/category"
	"github.com/tmazitov/tm_expenses_api/internal/domain/expense"
	"github.com/tmazitov/tm_expenses_api/internal/domain/user"
)

// Database

type mockDB struct{}

func (m mockDB) CategoryRepo() category.Repository { return mockCategoryRepo{} }
func (m mockDB) ExpenseRepo() expense.Repository   { return mockExpenseRepo{} }
func (m mockDB) UserRepo() user.Repository         { return mockUserRepo{} }

// Cache

type mockCache struct{}

func (m mockCache) Set(ctx context.Context, key, value string, ttl time.Duration) error { return nil }
func (m mockCache) Get(ctx context.Context, key string) (string, error)                 { return "", nil }
func (m mockCache) Del(ctx context.Context, keys ...string) error                       { return nil }

// Google Provider

type mockGoogleOAuthProvider struct{}

func (m mockGoogleOAuthProvider) ValidateToken(ctx context.Context, idToken string) (map[string]any, error) {
	return map[string]any{}, nil
}

// Jwt storage

type mockJwtProvider struct{}

func (m mockJwtProvider) CreateTokenPair(ctx context.Context, claims map[string]any) (string, string, error) {
	return "", "", nil
}
func (m mockJwtProvider) Refresh(ctx context.Context, token string) (string, string, error) {
	return "", "", nil
}
func (m mockJwtProvider) VerifyAccess(token string) (map[string]any, error) {
	return map[string]any{}, nil
}
func (m mockJwtProvider) VerifyRefresh(ctx context.Context, token string) (map[string]any, error) {
	return map[string]any{}, nil
}
func (m mockJwtProvider) Dispose(ctx context.Context, refreshToken string) error { return nil }
