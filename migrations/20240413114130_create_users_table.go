package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	tx.Exec(`
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	return nil
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	tx.Exec(`
		DROP TABLE users;
	`)
	return nil
}
