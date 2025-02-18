package test

import (
	"net/http"
	"net/http/httptest"
	"template/app"
	"template/store"
)

type TestCase struct {
	Store   *store.Queries
	Handler http.Handler
	Server  *httptest.Server
	Client  *http.Client
}

func NewTestCase() *TestCase {
	db := app.MustNewDatabase(":memory:")
	store := store.New(db)
	handler := app.NewHandler(store)
	server := httptest.NewServer(handler)

	return &TestCase{
		Store:   store,
		Handler: handler,
		Server:  server,
		Client:  server.Client(),
	}
}

func (tc *TestCase) Close() {
	tc.Server.Close()
}
