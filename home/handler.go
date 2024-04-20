package home

import (
	"workflou/template/auth"
	"workflou/template/templates/pages"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	AuthStore *auth.Store
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		AuthStore: auth.NewStore(db),
	}
}

func (h *Handler) HomePage(c echo.Context) error {
	user, err := h.AuthStore.FindByEmail(c.Request().Context(), "test@example.com")
	if err != nil {
		return c.String(404, "Not Found")
	}

	return pages.Home(pages.HomeDTO{User: user}).Render(c.Request().Context(), c.Response().Writer)
}
