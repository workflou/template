package main

import (
	"database/sql"
	"template/migrations"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func mustNewDatabase(dsn string) *sql.DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	goose.SetBaseFS(migrations.FS)

	if err = goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err = goose.Up(db, "."); err != nil {
		panic(err)
	}

	return db
}
