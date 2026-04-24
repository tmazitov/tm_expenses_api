package user

import "context"

type Repository interface {
	GetBySub(ctx context.Context, method AuthMethod, sub string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) error
}
