package user

import (
	"context"
	"errors"

	"github.com/tmazitov/tm_expenses_api/internal/domain/user"
)

func (r *Repository) GetById(ctx context.Context, userId int) (*user.User, error) {

	var UserModel *UserModel = &UserModel{}

	err := r.db.NewSelect().
		Model(UserModel).
		Where("id=?", userId).
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	user, err := user.NewUser(UserModel.ToUserParams())
	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	return user, nil
}
