package main

import (
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
)

var (
	//go:embed ui/dist/*
	uiFS embed.FS

	dsn = flag.String("dsn", "postgres://postgres:postgres@localhost:5432/postgres", "database connection string")
)

func main() {
	flag.Parse()

	// db := mustNewDatabase(*dsn)

	middleware := middlewareStack(recoverMiddleware, loggingMiddleware)

	router := http.NewServeMux()

	dist, _ := fs.Sub(uiFS, "ui/dist")
	router.Handle("/", http.FileServer(http.FS(dist)))

	v1 := http.NewServeMux()
	v1.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "hello, world"})
	}))
	router.Handle("/v1", authMiddleware(v1))

	s := http.Server{
		Addr:    ":4000",
		Handler: middleware(router),
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
