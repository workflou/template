package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"template/app"
	"template/store"
)

var (
	dsn = flag.String("dsn", ":memory:", "database connection string")
)

func main() {
	flag.Parse()

	db := app.MustNewDatabase(*dsn)
	handler := app.NewHandler(store.New(db))

	s := http.Server{
		Addr:    ":4000",
		Handler: handler,
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
