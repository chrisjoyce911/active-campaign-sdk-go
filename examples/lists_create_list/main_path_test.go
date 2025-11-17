package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain_CreateAndDeleteList(t *testing.T) {
	var createdID string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"list":{"id":"L1","name":"Example List"}}`)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/contactLists":
			// best-effort subscribe; return 200 OK JSON
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contactList":{"id":"cl1"}}`)
			return
		case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/api/3/lists/"):
			createdID = strings.TrimPrefix(r.URL.Path, "/api/3/lists/")
			w.WriteHeader(200)
			_, _ = io.WriteString(w, `{}`)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldSafe := os.Getenv("LISTS_SAFE")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("LISTS_SAFE", oldSafe)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("LISTS_SAFE", "false") // ensure delete flag defaults to true

	main()

	if createdID != "L1" {
		t.Fatalf("expected delete of L1, got %s", createdID)
	}
}
