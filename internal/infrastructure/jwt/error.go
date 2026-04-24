package jwt

import "errors"

var (
	ErrInvalidSecret        = errors.New("jwt storage error : secret parameter is invalid")
	ErrInvalidCacheProvider = errors.New("jwt storage error : cache provider parameter is nil")
	ErrInvalidTTL           = errors.New("jwt storage error : tokens' ttl is empty")

	ErrCreateAccessFailed  = errors.New("jwt storage error : failed to create an access token")
	ErrCreateRefreshFailed = errors.New("jwt storage error : failed to create an refresh token")
	ErrStoreTokenFailed    = errors.New("jwt storage error : failed to store a token in the cache.")

	ErrInvalidToken = errors.New("jwt storage error : token is invalid.")
)
