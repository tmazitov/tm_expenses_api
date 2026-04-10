package category

import "fmt"

type CategoryError struct {
	code    string
	message string
}

func NewCategoryError(code, message string) *CategoryError {
	return &CategoryError{
		code:    code,
		message: message,
	}
}

func (e *CategoryError) Error() string {
	return fmt.Sprintf("expense domain error [%s] : %s", e.code, e.message)
}

var (
	ErrIdEmpty     = NewCategoryError("NIL_ID", "category: id cannot be empty")
	ErrNameEmpty   = NewCategoryError("NIL_NAME", "category: id cannot be empty")
	ErrNameTooLong = NewCategoryError("NAME_TOO_LONG", "category: name is more than 100 symbols")
	ErrColorInvalidFormat = NewCategoryError("COLOR_INVALID_FORMAT", "category: color isn't hex format")
)
