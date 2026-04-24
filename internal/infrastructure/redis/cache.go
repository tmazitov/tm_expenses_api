package redis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

// redis.CacheClient is a client allows to save temporal data using Redis external service.
type CacheClient struct {
	redisClient *redis.Client
}

type CacheParams struct {
	Addr string
	DB   int
}

func (p CacheParams) validate() error {

	if p.DB < 0 || p.DB > 15 {
		return ErrInvalidDB
	}

	if _, err := net.ResolveTCPAddr("tcp", p.Addr); err != nil {
		return ErrInvalidAddr
	}

	return nil
}

// NewCacheClient creates a client using specific parameters.
func NewCacheClient(params CacheParams) (*CacheClient, error) {

	if err := params.validate(); err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         params.Addr,
		DB:           params.DB,
		MinIdleConns: 2,
	})

	if status := redisClient.Ping(context.Background()); status.Err() != nil {
		return nil, fmt.Errorf("redis cache client error : %w", status.Err())
	}

	return &CacheClient{
		redisClient: redisClient,
	}, nil
}

func (c *CacheClient) Close() error {
	return c.redisClient.Close()
}

// Sets value with particular key and ttl.
func (c *CacheClient) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	err := c.redisClient.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return errors.Join(ErrSetOperationFailed, err)
	}
	return nil
}

// Get returns record's value by key. If it doesn't exist, returns empty string.
func (c *CacheClient) Get(ctx context.Context, key string) (string, error) {
	status := c.redisClient.Get(ctx, key)
	err := status.Err()
	if err != nil && err != redis.Nil {
		return "", errors.Join(ErrGetOperationFailed, err)
	}
	return status.Val(), nil
}

// Del removes values by proceeded keys.
func (c *CacheClient) Del(ctx context.Context, keys ...string) error {
	err := c.redisClient.Del(ctx, keys...).Err()
	if err != nil {
		return errors.Join(ErrDelOperationFailed, err)
	}
	return nil
}
