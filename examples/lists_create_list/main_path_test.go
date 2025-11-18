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

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldSafe := os.Getenv("LISTS_SAFE")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("LISTS_SAFE", oldSafe)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("LISTS_SAFE", "false") // ensure delete flag defaults to true

	main()

	if createdID != "L1" {
		t.Fatalf("expected delete of L1, got %s", createdID)
	}
}

func TestMain_NoDelete(t *testing.T) {
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
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldSafe := os.Getenv("LISTS_SAFE")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("LISTS_SAFE", oldSafe)
	})
	os.Args = []string{"main", "-delete=false"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("LISTS_SAFE", "true")

	main()
}

func TestMain_UsesSearchByEmail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/3/contacts":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contacts":[{"id":"c-from-search"}]}`)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"list":{"id":"L2","name":"Example List"}}`)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/contactLists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contactList":{"id":"cl2"}}`)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldEmail := os.Getenv("ACTIVE_EMAIL")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("ACTIVE_EMAIL", oldEmail)
	})
	os.Args = []string{"main", "-delete=false"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "")
	_ = os.Setenv("ACTIVE_EMAIL", "search@example.com")

	main()
}

func TestMain_CreateListError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/3/lists" {
			http.Error(w, "error", 500)
			return
		}
		http.NotFound(w, r)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()
}

func TestMain_UpdateListStatusError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"list":{"id":"L3","name":"Example List"}}`)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/contactLists":
			http.Error(w, "error", 500)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
	})
	os.Args = []string{"main", "-delete=false"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")

	main()
}

func TestMain_DeleteError(t *testing.T) {
	var createdID string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"list":{"id":"L4","name":"Example List"}}`)
			return
		case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/api/3/lists/"):
			createdID = strings.TrimPrefix(r.URL.Path, "/api/3/lists/")
			http.Error(w, "error", 500)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	os.Args = []string{"main", "-delete=true"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()

	if createdID != "L4" {
		t.Fatalf("expected delete attempt for L4, got %s", createdID)
	}
}

func TestMain_SearchByEmailError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/3/contacts":
			http.Error(w, "error", 500)
			return
		case r.Method == http.MethodPost && r.URL.Path == "/api/3/lists":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"list":{"id":"L5","name":"Example List"}}`)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldEmail := os.Getenv("ACTIVE_EMAIL")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("ACTIVE_EMAIL", oldEmail)
	})
	os.Args = []string{"main", "-delete=false"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "")
	_ = os.Setenv("ACTIVE_EMAIL", "search@example.com")

	main()
}

func TestMain_MissingEnv(t *testing.T) {
	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	os.Args = []string{"main"}
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")
	main()
}

func TestMain_BadURL(t *testing.T) {
	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	main()
}
