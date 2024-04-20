package main

import (
	"workflou/template/db"
	"workflou/template/home"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := db.New()
	e := newServer()

	home := home.NewHandler(db)
	e.GET("/", home.HomePage)

	e.Logger.Fatal(e.Start(":4000"))
}
