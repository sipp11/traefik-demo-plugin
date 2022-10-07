package demoplugin_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	demoplugin "github.com/sipp11/traefik-demo-plugin"
)

func TestDemo(t *testing.T) {
	cfg := demoplugin.CreateConfig()
	cfg.Headers["X-Host"] = "[[.Host]]"
	cfg.Headers["X-Method"] = "[[.Method]]"
	cfg.Headers["X-URL"] = "[[.URL]]"
	cfg.Headers["X-URLX"] = "[[.URL]]"
	cfg.Headers["X-Demo-One"] = "test"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := demoplugin.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Host", "localhost")
	assertHeader(t, req, "X-URL", "http://localhost")
	assertHeader(t, req, "X-URLX", "http://localhost")
	assertHeader(t, req, "X-Method", "GET")
	assertHeader(t, req, "X-Demo-One", "test")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
