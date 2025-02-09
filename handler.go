package main

import (
	"net/http"
	"template/store"
	"template/view"
)

type handler struct {
	Store *store.Queries
}

func (h *handler) HomePage(w http.ResponseWriter, r *http.Request) {
	view.HomePage().Render(r.Context(), w)
}

func (h *handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	view.LoginPage().Render(r.Context(), w)
}
