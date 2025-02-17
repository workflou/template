package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"template/static"
	"template/store"
)

var (
	dsn = flag.String("dsn", ":memory:", "database connection string")
)

func main() {
	flag.Parse()

	db := mustNewDatabase(*dsn)

	handler := &handler{
		Store: store.New(db),
	}
	middleware := middlewareStack(recoverMiddleware, loggingMiddleware)

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))
	router.HandleFunc("/login", handler.LoginPage)

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/{$}", handler.HomePage)

	router.Handle("/", authMiddleware(userRouter))

	s := http.Server{
		Addr:    ":4000",
		Handler: middleware(router),
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
