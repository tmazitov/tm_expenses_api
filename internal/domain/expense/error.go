package expense

import (
	"fmt"
)

type ExpenseError struct {
	code    string
	message string
}

func NewExpenseError(code, message string) *ExpenseError {
	return &ExpenseError{
		code:    code,
		message: message,
	}
}

func (e *ExpenseError) Error() string {
	return fmt.Sprintf("expense domain error [%s] : %s", e.code, e.message)
}

var (
	ErrEmptyName   = NewExpenseError("EMPTY_NAME", "expense: name cannot be empty")
	ErrNameTooLong = NewExpenseError("NAME_TOO_LONG", "expense: name cannot exceed 100 characters")
)
