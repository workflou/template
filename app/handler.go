package app

import (
	"net/http"
	"template/static"
	"template/store"
	"template/view"
)

type Handler struct {
	Store *store.Queries
}

func NewHandler(Store *store.Queries) http.Handler {
	handler := &Handler{Store: Store}

	middleware := NewMiddlewareStack(recoverMiddleware, loggingMiddleware)

	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))
	router.HandleFunc("/login", handler.LoginPage)

	{
		userRouter := http.NewServeMux()
		userRouter.HandleFunc("/{$}", handler.HomePage)

		router.Handle("/", authMiddleware(userRouter))
	}

	return middleware(router)
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	view.HomePage().Render(r.Context(), w)
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	view.LoginPage().Render(r.Context(), w)
}
