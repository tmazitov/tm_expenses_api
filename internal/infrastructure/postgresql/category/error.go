package category

import "errors"

var (
	ErrInsertionFailed = errors.New("category repo error: insertion failed")
	ErrSelectionFailed = errors.New("category repo error: selection failed")
)
