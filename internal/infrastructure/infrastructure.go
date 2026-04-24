package infra

import (
	"fmt"

	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/google"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/jwt"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/redis"
)

type Infrastructure struct {
	db          *postgresql.Database
	cache       *redis.CacheClient
	googleOAuth *google.OAuthProvider
	jwt         *jwt.Storage
}

type InfrastructureParams struct {
	DBConfig          postgresql.Config
	CacheParams       redis.CacheParams
	GoogleOAuthParams google.OAuthProviderParams
	JwtParams         jwt.StorageParams
}

func NewInfrastructure(params InfrastructureParams) (*Infrastructure, error) {

	db, err := postgresql.NewDatabase("./db/migrations", params.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("infrastructure setup error : %w", err)
	}

	cache, err := redis.NewCacheClient(params.CacheParams)
	if err != nil {
		return nil, fmt.Errorf("infrastructure setup error : %w", err)
	}

	googleOAuth, err := google.NewOAuthProvider(params.GoogleOAuthParams)
	if err != nil {
		return nil, fmt.Errorf("infrastructure setup error : %w", err)
	}

	jwtStorage, err := jwt.NewStorage(cache, params.JwtParams)
	if err != nil {
		return nil, fmt.Errorf("infrastructure setup error : %w", err)
	}

	return &Infrastructure{
		db:          db,
		cache:       cache,
		googleOAuth: googleOAuth,
		jwt:         jwtStorage,
	}, nil
}

func (i *Infrastructure) DB() *postgresql.Database           { return i.db }
func (i *Infrastructure) Cache() *redis.CacheClient          { return i.cache }
func (i *Infrastructure) Jwt() *jwt.Storage                  { return i.jwt }
func (i *Infrastructure) GoogleOAuth() *google.OAuthProvider { return i.googleOAuth }

func (i *Infrastructure) Close() {
	i.db.Close()
	i.cache.Close()
}
