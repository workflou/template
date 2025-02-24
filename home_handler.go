package main

import (
	"net/http"
	"template/view"
)

type HomeHandler struct {
}

func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	view.HomePage().Render(r.Context(), w)
}
