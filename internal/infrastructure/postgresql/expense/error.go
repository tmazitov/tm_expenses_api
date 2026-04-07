package expense

import "errors"

var (
	ErrInsertionFailed = errors.New("expense repo error: insertion failed")
	ErrSelectionFailed = errors.New("expense repo error: selection failed")
)
