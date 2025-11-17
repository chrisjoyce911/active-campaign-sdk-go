package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain_DryRunWithEnvContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/contacts" {
			w.Header().Set("Content-Type", "application/json")
			// return no matches to exercise env-provided contact id path
			_, _ = io.WriteString(w, `{"contacts":[]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer ts.Close()

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("CONTACT_EMAIL")
	oldCID := os.Getenv("ACTIVE_CONTACT_ID")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("CONTACT_EMAIL", oldEmail)
		_ = os.Setenv("ACTIVE_CONTACT_ID", oldCID)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("CONTACT_EMAIL", "test@example.com")
	_ = os.Setenv("ACTIVE_CONTACT_ID", "c1")

	main()
}
