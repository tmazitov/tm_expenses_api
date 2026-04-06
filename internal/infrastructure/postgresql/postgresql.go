package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	expenseDomain "github.com/tmazitov/ayda-order-service.git/internal/domain/expense"
	"github.com/tmazitov/ayda-order-service.git/internal/infrastructure/postgresql/expense"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

type Database struct {
	expenseRepository *expense.Repository
	instance          *bun.DB
}

func NewDatabase(migrationsDir string, config Config) (*Database, error) {

	sqlDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(config.DSN()),
	))
	instance := bun.NewDB(sqlDB, pgdialect.New())

	if err := instance.PingContext(context.Background()); err != nil {
		return nil, errors.Join(ErrConnectionFailed, err)
	}

	db := &Database{
		instance:          instance,
		expenseRepository: expense.NewRepository(instance),
	}

	if err := db.syncMigrations(migrationsDir); err != nil {
		return nil, errors.Join(ErrSyncMigrationsFailed, err)
	}

	return db, nil
}

func (db *Database) Close() error {
	return db.instance.Close()
}

func (db *Database) ExpenseRepo() expenseDomain.Repository { return db.expenseRepository }
