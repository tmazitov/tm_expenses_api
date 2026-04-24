package user

import (
	"time"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users"`

	Id         int             `bun:"id,pk"`
	FirstName  string          `bun:"first_name,notnull"`
	LastName   *string         `bun:"last_name,default:null"`
	Email      string          `bun:"email,notnull"`
	Sub        string          `bun:"sub,notnull"`
	AuthMethod user.AuthMethod `bun:"auth_method,notnull"`
	CreatedAt  *time.Time      `bun:"created_at,default:now()"`
}

func newUserModel(u *user.User) *UserModel {

	model := UserModel{
		Id:         u.Id(),
		FirstName:  u.FirstName(),
		Email:      u.Email(),
		AuthMethod: u.AuthMethod(),
		Sub:        u.Sub(),
	}

	createdAt := u.CreatedAt()
	if !createdAt.IsZero() {
		model.CreatedAt = &createdAt
	}
	lastName := u.LastName()
	if len(lastName) != 0 {
		model.LastName = &lastName
	}

	return &model
}

func (m *UserModel) ToUserParams() user.UserParams {

	params := user.UserParams{
		Id:         m.Id,
		FirstName:  m.FirstName,
		Email:      m.Email,
		AuthMethod: m.AuthMethod,
		Sub:        m.Sub,
	}

	if m.LastName != nil {
		params.LastName = *m.LastName
	}

	if m.CreatedAt != nil {
		params.CreatedAt = *m.CreatedAt
	}

	return params
}
