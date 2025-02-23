package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

type Middleware func(http.Handler) http.Handler

// todo: add "new" prefix

func NewMiddlewareStack(m ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			h = m[i](h)
		}

		return h
	}
}

func NewRecoverMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			h.ServeHTTP(w, r)
		})
	}
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func NewLoggingMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()

			rw := &wrappedResponseWriter{w, http.StatusOK}

			h.ServeHTTP(rw, r)

			slog.Info(fmt.Sprintf("%d %s %s %s", rw.status, r.Method, r.URL.Path, time.Since(t)))
		})
	}
}

func NewAuthMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			if c.Value == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
