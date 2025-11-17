package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestMain_HappyPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// SearchContacts via /contacts?email=...
		if r.URL.Path == "/api/3/contacts" {
			if q, _ := url.QueryUnescape(r.URL.RawQuery); q == "" {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contacts":[{"id":"c1"}]}`)
			return
		}
		// GetAutomationCounts via /contactAutomations/counts
		if r.URL.Path == "/api/3/contactAutomations/counts" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"counts":[{"automation":"a1","count":1}]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer ts.Close()

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("ACTIVE_EMAIL")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_EMAIL", oldEmail)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_EMAIL", "a@b.com")

	main()
}
