package main

import (
	"net/http"
	"template/store"
	"template/view"
)

type AccountHandler struct {
	Store *store.Queries
}

func (h *AccountHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	view.LoginPage().Render(r.Context(), w)
}
