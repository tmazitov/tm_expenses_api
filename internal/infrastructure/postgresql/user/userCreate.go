package user

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

func (r *Repository) Create(ctx context.Context, u *user.User) (*user.User, error) {

	model := newUserModel(u)

	_, err := r.db.NewInsert().
		Model(model).
		Exec(ctx)

	if err != nil {
		return nil, errors.Join(ErrInsertionFailed, err)
	}

	u, err = user.NewUser(model.ToUserParams())
	if err != nil {
		return nil, errors.Join(ErrInsertionFailed, err)
	}

	return u, nil
}
