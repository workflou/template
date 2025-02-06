package main

import (
	"log"
	"log/slog"
	"net/http"
	"template/static"
)

func main() {
	h := &handler{}
	m := middlewareStack(recoverMiddleware, loggingMiddleware)
	r := http.NewServeMux()

	r.HandleFunc("/{$}", h.Home())
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))

	s := http.Server{
		Addr:    ":4000",
		Handler: m(r),
	}

	slog.Info("http://localhost:4000")
	log.Fatal(s.ListenAndServe())
}
