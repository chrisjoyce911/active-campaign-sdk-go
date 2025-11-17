package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain_Generate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/tags":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"tags":[{"id":"1","tag":"VIP"}]}`)
		case "/api/3/fields":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"fields":[{"id":"2","title":"Company"}]}`)
		case "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"lists":[{"id":"3","name":"Newsletter"}]}`)
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

func TestMain_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server error", 500)
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

