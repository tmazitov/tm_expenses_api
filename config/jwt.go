package config

import (
	"errors"
	"strconv"
	"time"
)

type JWT struct {
	Secret     []byte
	AccessTTL  time.Duration // minutes
	RefreshTTL time.Duration // days
}

func NewJWT() (*JWT, error) {

	secret := getenvDefault("JWT_SECRET", "")
	if len(secret) == 0 {
		return nil, errors.New("jwt config error : secret is empty")
	}

	accessTTLRaw := getenvDefault("JWT_ACCESS_TTL", "15")
	accessTTL, err := strconv.Atoi(accessTTLRaw)
	if err != nil {
		return nil, errors.New("jwt config error : access ttl is invalid")
	}

	refreshTTLRaw := getenvDefault("JWT_REFRESH_TTL", "15")
	refreshTTL, err := strconv.Atoi(refreshTTLRaw)
	if err != nil {
		return nil, errors.New("jwt config error : refresh ttl is invalid")
	}

	return &JWT{
		Secret:     []byte(secret),
		AccessTTL:  time.Minute * time.Duration(accessTTL),
		RefreshTTL: time.Minute * time.Duration(refreshTTL),
	}, nil
}
