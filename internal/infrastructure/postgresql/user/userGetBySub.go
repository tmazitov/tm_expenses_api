package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

func (r *Repository) GetBySub(ctx context.Context, method user.AuthMethod, sub string) (*user.User, error) {

	var model *UserModel = &UserModel{}

	err := r.db.NewSelect().
		Model(model).
		Where("auth_method=?", method).
		Where("sub=?", sub).
		Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	u, err := user.NewUser(model.ToUserParams())
	if err != nil {
		return nil, errors.Join(ErrSelectionFailed, err)
	}

	return u, nil
}
