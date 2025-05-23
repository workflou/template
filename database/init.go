package database

import (
	"database/sql"
	"embed"
	"io/fs"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var (
	DB *sql.DB

	//go:embed schema/*.sql
	FS embed.FS
)

func init() {
	var err error

	DB, err = sql.Open("sqlite", os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	schemaFS, err := fs.Sub(FS, "schema")
	if err != nil {
		panic(err)
	}

	goose.SetDialect("sqlite3")
	goose.SetBaseFS(schemaFS)

	if err = goose.Up(DB, "."); err != nil {
		panic(err)
	}
}
