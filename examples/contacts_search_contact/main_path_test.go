package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain_HappyPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/3/contacts" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"contacts":[{"id":"c1"}]}`)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("CONTACT_EMAIL")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("CONTACT_EMAIL", oldEmail)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("CONTACT_EMAIL", "a@b.com")

	main()
}

func TestMain_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server error", 500)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("CONTACT_EMAIL")
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("CONTACT_EMAIL", oldEmail)
		_ = os.Setenv("TEST", oldTest)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("CONTACT_EMAIL", "a@b.com")
	_ = os.Setenv("TEST", "1")

	main()
}
