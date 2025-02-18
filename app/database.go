package app

import (
	"database/sql"
	"template/schema"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func MustNewDatabase(dsn string) *sql.DB {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	goose.SetBaseFS(schema.FS)

	if err = goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err = goose.Up(db, "."); err != nil {
		panic(err)
	}

	return db
}
