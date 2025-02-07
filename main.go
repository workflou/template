package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"template/static"
)

var (
	dsn = flag.String("dsn", "postgres://postgres:postgres@localhost:5432/postgres", "database connection string")
)

func main() {
	flag.Parse()

	db := mustNewDatabase(*dsn)
	handler := &handler{
		DB: db,
	}
	middleware := middlewareStack(recoverMiddleware, loggingMiddleware)

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))
	router.HandleFunc("/login", handler.LoginPage())

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("/{$}", handler.Home())

	router.Handle("/", authMiddleware(authRouter))

	s := http.Server{
		Addr:    ":4000",
		Handler: middleware(router),
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
