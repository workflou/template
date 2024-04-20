package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() *sqlx.DB {
	db := sqlx.MustConnect("postgres", os.Getenv("DATABASE_URL"))
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
