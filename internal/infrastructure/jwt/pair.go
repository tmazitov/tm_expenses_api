package jwt

import "time"

type tokenPair struct {
	Access  string
	Refresh string
}

type PairTTL struct {
	Access  time.Duration
	Refresh time.Duration
}
