package postgresql

import (
	"log"

	"github.com/pressly/goose/v3"
)

func (db *Database) syncMigrations(migrationsDir string) error {
	var err error

	goose.SetLogger(log.Default())

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(db.instance.DB, migrationsDir); err != nil {
		return err
	}
	return nil
}
