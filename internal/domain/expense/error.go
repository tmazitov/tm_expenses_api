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
	ErrPriceIsNegative = NewExpenseError("NEG_PRICE", "expense: price cannot be negative")

	ErrEmptyName   = NewExpenseError("EMPTY_NAME", "expense: name cannot be empty")
	ErrNameTooLong = NewExpenseError("NAME_TOO_LONG", "expense: name cannot exceed 255 characters")

	ErrPageIsNegative  = NewExpenseError("NEG_PAGE", "expense list: page cannot be negative")
	ErrLimitIsNegative = NewExpenseError("NEG_LIMIT", "expense list: limit cannot be negative")
)
