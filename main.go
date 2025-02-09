package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"template/static"
	"template/store"
)

var (
	dsn = flag.String("dsn", "postgres://postgres:postgres@localhost:5432/postgres", "database connection string")
)

func main() {
	flag.Parse()

	db := mustNewDatabase(*dsn)
	st := store.New(db)

	_, err := st.GetUserByID(context.Background(), 1)
	if err != nil {
		st.CreateUser(context.Background(), store.CreateUserParams{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password", // todo
		})
	}

	handler := &handler{
		Store: st,
	}
	middleware := middlewareStack(recoverMiddleware, loggingMiddleware)

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))
	router.HandleFunc("/login", handler.LoginPage)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("/{$}", handler.HomePage)

	router.Handle("/", authMiddleware(authRouter))

	s := http.Server{
		Addr:    ":4000",
		Handler: middleware(router),
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
