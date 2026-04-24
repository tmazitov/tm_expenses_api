package user

import "errors"

var (
	ErrInvalidParams      error = errors.New("user service : invalid constructor params")
	ErrInvalidCredentials error = errors.New("user service : received credentials to authorize user is invalid")
	ErrGoogleOauthFailed  error = errors.New("user service : google oauth process failed")
)
