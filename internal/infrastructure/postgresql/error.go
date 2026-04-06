package postgresql

import "errors"

var (
	ErrConnectionFailed     = errors.New("postgresql db error: connection failed")
	ErrSyncMigrationsFailed = errors.New("postgresql db error: sync migration process failed")
)
