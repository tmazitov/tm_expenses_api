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

	ErrPageIsNegative  = NewExpenseError("NEG_PAGE", "expense filters: page cannot be negative")
	ErrLimitIsNegative = NewExpenseError("NEG_LIMIT", "expense filters: limit cannot be negative")
	ErrDateIsZero      = NewExpenseError("ZERO_DATE", "expense filters: date cannot be zero")
)

// ExpenseStat
var (
	ErrInvalidWeeklyID    = NewExpenseError("INVALID_WEEKLY_ID", "expense stat identifier (weekly): value cannot exceed 52")
	ErrInvalidMonthlyID   = NewExpenseError("INVALID_MONTHLY_ID", "expense stat identifier (monthly): value cannot exceed 12")
	ErrUnknownStatVariant = NewExpenseError("UNKNOWN_STAT", "expense stat: variant is unknown")
)

// ExpenseStatFilters
var (
	ErrUnitsIsNegative = NewExpenseError("NEG_LIMIT", "expense filters: units cannot exceed 6")
)
