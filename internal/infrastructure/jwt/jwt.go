package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CacheProvider interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
}

type Storage struct {
	secret     []byte
	cache      CacheProvider
	accessTTL  time.Duration
	refreshTTL time.Duration
}

type StorageParams struct {
	Secret     []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func (p *StorageParams) validate() error {
	if len(p.Secret) == 0 {
		return ErrInvalidSecret
	}

	if p.AccessTTL <= 0 || p.RefreshTTL <= 0 {
		return ErrInvalidTTL
	}
	return nil
}

func NewStorage(cache CacheProvider, params StorageParams) (*Storage, error) {
	if cache == nil {
		return nil, ErrInvalidCacheProvider
	}
	if err := params.validate(); err != nil {
		return nil, err
	}

	return &Storage{
		cache:      cache,
		secret:     params.Secret,
		accessTTL:  params.AccessTTL,
		refreshTTL: params.RefreshTTL,
	}, nil
}

// CreateTokenPair creates a new pair of jwt tokens for the user.
// As additional information, it receives claims and TTL for tokens.
// Proceeded keys "exp" and "jti" will be ignored in claims. They will be set up automatically.
func (s *Storage) CreateTokenPair(ctx context.Context, claims map[string]any) (string, string, error) {
	accessClaims := jwt.MapClaims{"exp": time.Now().Add(s.accessTTL).Unix(), "jti": uuid.NewString()}
	refreshClaims := jwt.MapClaims{"exp": time.Now().Add(s.refreshTTL).Unix(), "jti": uuid.NewString()}

	for key, value := range claims {
		if key == "exp" || key == "jti" {
			continue
		}
		accessClaims[key] = value
		refreshClaims[key] = value
	}

	access, err := s.newToken(accessClaims)
	if err != nil {
		return "", "", errors.Join(ErrCreateAccessFailed, err)
	}

	refresh, err := s.newToken(refreshClaims)
	if err != nil {
		return "", "", errors.Join(ErrCreateRefreshFailed, err)
	}

	// Save refresh token to the cache
	key := fmt.Sprintf("refresh:%s", refreshClaims["jti"])
	if err := s.cache.Set(ctx, key, refresh, s.refreshTTL); err != nil {
		return "", "", errors.Join(ErrStoreTokenFailed, err)
	}

	return access, refresh, nil
}

func (s *Storage) Refresh(ctx context.Context, token string) (string, string, error) {

	claims, err := s.VerifyRefresh(ctx, token)
	if err != nil {
		return "", "", err
	}

	jti, ok := claims["jti"]
	if !ok {
		return "", "", ErrInvalidToken
	}

	key := fmt.Sprintf("refresh:%s", jti)
	if err := s.cache.Del(ctx, key); err != nil {
		return "", "", errors.Join(ErrInvalidToken, err)
	}

	return s.CreateTokenPair(ctx, claims)
}

func (s *Storage) VerifyAccess(token string) (map[string]any, error) {
	return s.verifyToken(token)
}

func (s *Storage) VerifyRefresh(ctx context.Context, token string) (map[string]any, error) {
	claims, err := s.verifyToken(token)
	if err != nil {
		return nil, err
	}

	jti := claims["jti"]

	key := fmt.Sprintf("refresh:%s", jti)
	storedToken, err := s.cache.Get(ctx, key)
	if err != nil {
		return nil, errors.Join(ErrInvalidToken, err)
	}

	if storedToken != token {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (s *Storage) verifyToken(token string) (jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, s.verifier)
	if err != nil {
		return nil, errors.Join(ErrInvalidToken, err)
	}
	return *claims, nil
}

func (s *Storage) newToken(claims jwt.MapClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secret)
}

func (s *Storage) verifier(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(s.secret), nil
}
