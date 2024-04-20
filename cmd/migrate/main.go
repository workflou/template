package main

import (
	"workflou/template/db"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "workflou/template/migrations"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := db.New()
	goose.SetDialect("postgres")

	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}
}
