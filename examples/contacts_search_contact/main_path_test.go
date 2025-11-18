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

func TestMain_MissingArgs_WithExit(t *testing.T) {
	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	oldExit := exitFn
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
		exitFn = oldExit
	})
	code := 0
	exitFn = func(c int) { code = c }
	os.Args = []string{"main"}
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")
	_ = os.Unsetenv("CONTACT_EMAIL")
	_ = os.Setenv("TEST", "2")
	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_BadURL_WithExit(t *testing.T) {
	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	oldExit := exitFn
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
		exitFn = oldExit
	})
	code := 0
	exitFn = func(c int) { code = c }
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("CONTACT_EMAIL", "a@b.com")
	_ = os.Setenv("TEST", "2")
	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_APIError_WithExit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server error", 500)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	oldExit := exitFn
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
		exitFn = oldExit
	})
	code := 0
	exitFn = func(c int) { code = c }
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("CONTACT_EMAIL", "a@b.com")
	_ = os.Setenv("TEST", "2")
	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_MissingArgs(t *testing.T) {
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
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")
	_ = os.Unsetenv("CONTACT_EMAIL")
	_ = os.Setenv("TEST", "1")

	main()
}

func TestMain_BadURL(t *testing.T) {
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
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("CONTACT_EMAIL", "a@b.com")
	_ = os.Setenv("TEST", "1")

	main()
}
