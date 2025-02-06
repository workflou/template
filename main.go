package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	m := middlewareStack(recoverMiddleware, loggingMiddleware)

	r := http.NewServeMux()

	r.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	s := http.Server{
		Addr:    ":4000",
		Handler: m(r),
	}

	slog.Info("http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
