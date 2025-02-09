package main

import (
	"database/sql"
	"net/http"
	"template/view"
)

type handler struct {
	DB *sql.DB
}

func (h *handler) HomePage(w http.ResponseWriter, r *http.Request) {
	view.HomePage().Render(r.Context(), w)
}

func (h *handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	view.LoginPage().Render(r.Context(), w)
}
