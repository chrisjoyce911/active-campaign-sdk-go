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
		if r.URL.Path != "/api/3/campaigns" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"campaigns":[{"id":"c1","name":"One","status":"1"}],"meta":{"total":"1"}}`)
	}))
	defer ts.Close()

	// Set env expected by main
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	// Call main; should not exit on success
	main()
}

func TestMain_MissingEnv_ShouldExit(t *testing.T) {
	oldExit := exitFn
	defer func() { exitFn = oldExit }()
	code := 0
	exitFn = func(c int) { code = c }

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")

	main()
	if code == 0 {
		t.Fatalf("expected exit code > 0 for missing env")
	}
}

func TestMain_BadURL_ShouldExit(t *testing.T) {
	oldExit := exitFn
	defer func() { exitFn = oldExit }()
	code := 0
	exitFn = func(c int) { code = c }

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()
	if code == 0 {
		t.Fatalf("expected exit code > 0 for bad url")
	}
}

func TestMain_RunError_ShouldExit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server error", 500)
	}))
	defer ts.Close()

	oldExit := exitFn
	defer func() { exitFn = oldExit }()
	code := 0
	exitFn = func(c int) { code = c }

	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
	})
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")

	main()
	if code == 0 {
		t.Fatalf("expected exit code > 0 for run error")
	}
}
