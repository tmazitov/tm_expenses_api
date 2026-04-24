package redis

import "errors"

var (
	ErrInvalidDB   = errors.New("redis cache client error : db argument must be in range from 0 to 15.")
	ErrInvalidAddr = errors.New("redis cache client error : addr argument has invalid format.")

	ErrSetOperationFailed = errors.New("redis cache client error : set operation failed")
	ErrGetOperationFailed = errors.New("redis cache client error : get operation failed")
	ErrDelOperationFailed = errors.New("redis cache client error : del operation failed")
)
