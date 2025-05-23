package main

import (
	"flag"
	"net/http"
	"template/html"
	"template/static"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	addr = flag.String("addr", ":4000", "HTTP server address")
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.NoCache)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, world!"))
		})
	})

	r.Get("/", templ.Handler(html.HomePage()).ServeHTTP)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))

	s := http.Server{
		Addr:    *addr,
		Handler: r,
	}

	s.ListenAndServe()
}
