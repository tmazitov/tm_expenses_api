package config

import (
	"fmt"
	"net"
	"strconv"
)

type Cache struct {
	Addr string
	DB   int
}

func NewCache() (*Cache, error) {

	addr := getenvDefault("CACHE_ADDR", "localhost:6379")
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		return nil, fmt.Errorf("cache config error : addr parameter should follow the format \"host:port\"")
	}

	dbRaw := getenvDefault("CACHE_DB", "0")
	db, err := strconv.Atoi(dbRaw)
	if err != nil || db < 0 || db > 15 {
		return nil, fmt.Errorf("cache config error : db parameter should be in range from 0 to 15")
	}

	return &Cache{
		Addr: addr,
		DB:   db,
	}, nil
}
