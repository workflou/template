package test

import "testing"

func TestLoginPageRedirect(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	res, err := tc.Client.Get(tc.Server.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.Request.Response.StatusCode != 303 {
		t.Fatalf("expected status code 303, got %d", res.StatusCode)
	}

	if location := res.Request.Response.Header.Get("Location"); location != "/login" {
		t.Fatalf("expected location /login, got %s", location)
	}
}
