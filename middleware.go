package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

type middleware func(http.Handler) http.Handler

func middlewareStack(m ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			h = m[i](h)
		}

		return h
	}
}

func recoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		rw := &wrappedResponseWriter{w, http.StatusOK}

		h.ServeHTTP(rw, r)

		slog.Info(fmt.Sprintf("%d %s %s %s", rw.status, r.Method, r.URL.Path, time.Since(t)))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}

		if c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
