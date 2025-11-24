package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

// Run applies all pending database migrations using goose
func Run(db *sql.DB) error {
	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}

	return nil
}
