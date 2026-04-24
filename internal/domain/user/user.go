package user

import (
	"net/mail"
	"time"
)

type User struct {
	id         int
	authMethod AuthMethod
	firstName  string
	lastName   string
	email      string
	sub        string
	createdAt  time.Time
}

type UserParams struct {
	Id         int
	AuthMethod AuthMethod
	FirstName  string
	LastName   string
	Email      string
	Sub        string
	CreatedAt  time.Time
}

func (p *UserParams) validate() error {

	if len(p.FirstName) == 0 {
		return ErrEmptyFirstName
	} else if len(p.FirstName) > 255 {
		return ErrTooLongFirstName
	}

	if len(p.LastName) > 255 {
		return ErrTooLongLastName
	}

	if len(p.Email) == 0 {
		return ErrEmptyEmail
	} else if len(p.Email) > 255 {
		return ErrTooLongEmail
	} else if _, err := mail.ParseAddress(p.Email); err != nil {
		return ErrInvalidFormatEmail
	}

	return nil
}

// NewUser is a constructor dedicated to user instance.
func NewUser(params UserParams) (*User, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &User{
		id:         params.Id,
		sub:        params.Sub,
		email:      params.Email,
		authMethod: params.AuthMethod,
		firstName:  params.FirstName,
		lastName:   params.LastName,
		createdAt:  params.CreatedAt,
	}, nil
}

func (u *User) WithUpdatedInfo(params UserParams) (*User, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	return &User{
		id:         u.id,
		sub:        u.sub,
		createdAt:  u.createdAt,
		authMethod: u.authMethod,
		firstName:  params.FirstName,
		lastName:   params.LastName,
		email:      params.Email,
	}, nil
}

func (u *User) Id() int                { return u.id }
func (u *User) FirstName() string      { return u.firstName }
func (u *User) LastName() string       { return u.lastName }
func (u *User) Email() string          { return u.email }
func (u *User) Sub() string            { return u.sub }
func (u *User) AuthMethod() AuthMethod { return u.authMethod }
func (u *User) CreatedAt() time.Time   { return u.createdAt }
