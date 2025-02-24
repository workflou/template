package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"template/static"
	"template/store"
)

func main() {
	dsn := flag.String("dsn", ":memory:", "database connection string")
	flag.Parse()

	db := MustNewDatabase(*dsn)

	middleware := NewMiddlewareStack(NewRecoverMiddleware(), NewLoggingMiddleware())

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))

	accountHandler := &AccountHandler{
		Store: store.New(db),
	}

	router.HandleFunc("/login", accountHandler.LoginPage)

	homeHandler := &HomeHandler{}

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/{$}", homeHandler.HomePage)

	router.Handle("/", NewAuthMiddleware()(userRouter))

	s := http.Server{
		Addr:    ":4000",
		Handler: middleware(router),
	}

	slog.Info("http://localhost:4000")

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	slog.Info("shutting down")
}
