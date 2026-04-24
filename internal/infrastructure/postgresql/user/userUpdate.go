package user

import (
	"context"
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

func (r *Repository) Update(ctx context.Context, u *user.User) error {

	model := newUserModel(u)

	_, err := r.db.NewUpdate().
		Table("users").
		Set("first_name=?", model.FirstName).
		Set("last_name=?", model.LastName).
		Set("email=?", model.Email).
		Where("id=?", u.Id()).
		Exec(ctx)

	if err != nil {
		return errors.Join(ErrUpdateFailed, err)
	}

	return nil
}
