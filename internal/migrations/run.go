package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func Run(db *sql.DB) error {
	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}

	return nil
}
