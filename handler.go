package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"template/html"
)

type handler struct {
	DB *sql.DB
}

func (h *handler) Home() http.HandlerFunc {
	t := template.Must(template.ParseFS(html.FS, "layout.html", "home.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
