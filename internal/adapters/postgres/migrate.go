package postgres

import (
	"database/sql"

	"github.com/pressly/goose"
)

func Migrate(
	conn *sql.DB,
	migrationPath string,
) error {
	return goose.Up(conn, migrationPath)
}
