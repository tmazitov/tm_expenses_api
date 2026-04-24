package user

import "errors"

var (
	ErrSelectionFailed = errors.New("user repo error: selection failed")
	ErrInsertionFailed = errors.New("user repo error: insertion failed")
	ErrUpdateFailed    = errors.New("user repo error: update failed")
)
