//go:build examples

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain_CreateContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contact":{"id":"123","email":"jdoe@example.com"}}`)
			}
		case "/api/3/contactTags":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contactTag":{"id":"456"}}`)
			}
		case "/api/3/fieldValues":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"fieldValue":{"id":"789"}}`)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()
}
