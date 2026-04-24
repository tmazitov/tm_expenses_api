package user

import "fmt"

type UserError struct {
	code    string
	message string
}

func NewUserError(code, message string) *UserError {
	return &UserError{
		code:    code,
		message: message,
	}
}

func (e *UserError) Error() string {
	return fmt.Sprintf("expense domain error [%s] : %s", e.code, e.message)
}

var (
	ErrEmptyFirstName = NewUserError("EMPTY_FIRST_NAME", "user: first name cannot be empty")
	ErrEmptyEmail     = NewUserError("EMPTY_EMAIL", "user: email cannot be empty")

	ErrTooLongFirstName = NewUserError("FIRST_NAME_TOO_LONG", "user: first name is more than 255 symbols")
	ErrTooLongLastName  = NewUserError("LAST_NAME_TOO_LONG", "user: last name is more than 255 symbols")
	ErrTooLongEmail     = NewUserError("EMAIL_TOO_LONG", "user: email is more than 255 symbols")

	ErrInvalidFormatEmail = NewUserError("EMAIL_INVALID_FORMAT", "user: email has an invalid format")
)
