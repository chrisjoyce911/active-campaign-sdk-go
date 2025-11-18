package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain_DryRun(t *testing.T) {
	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldCompany := os.Getenv("CONTACT_COMPANY_NAME")
	oldField := os.Getenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("CONTACT_COMPANY_NAME", oldCompany)
		_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", oldField)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", "http://example")
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")

	main()
}

func TestMain_ApplyModeError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", 400)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldCompany := os.Getenv("CONTACT_COMPANY_NAME")
	oldField := os.Getenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("CONTACT_COMPANY_NAME", oldCompany)
		_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", oldField)
		_ = os.Setenv("TEST", oldTest)
	})
	os.Args = []string{"main", "-apply"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")
	_ = os.Setenv("TEST", "1")

	main()
}

func TestMain_ApplyModeError_WithExit(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", 400)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldExit := exitFn
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		exitFn = oldExit
		_ = os.Setenv("TEST", oldTest)
	})
	code := 0
	exitFn = func(c int) { code = c }

	os.Args = []string{"main", "-apply"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")
	_ = os.Setenv("TEST", "2")

	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_MissingEnv_WithExit(t *testing.T) {
	oldArgs := os.Args
	oldExit := exitFn
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		exitFn = oldExit
		_ = os.Setenv("TEST", oldTest)
	})
	code := 0
	exitFn = func(c int) { code = c }

	os.Args = []string{"main"}
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")
	_ = os.Unsetenv("ACTIVE_CONTACTID")
	_ = os.Unsetenv("CONTACT_COMPANY_NAME")
	_ = os.Unsetenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	_ = os.Setenv("TEST", "2")
	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_BadURL_WithExit(t *testing.T) {
	oldArgs := os.Args
	oldExit := exitFn
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		exitFn = oldExit
		_ = os.Setenv("TEST", oldTest)
	})
	code := 0
	exitFn = func(c int) { code = c }

	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")
	_ = os.Setenv("TEST", "2")
	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}

func TestMain_MissingEnv(t *testing.T) {
	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
	})
	os.Args = []string{"main"}
	_ = os.Unsetenv("ACTIVE_URL")
	_ = os.Unsetenv("ACTIVE_TOKEN")
	_ = os.Unsetenv("ACTIVE_CONTACTID")
	_ = os.Unsetenv("CONTACT_COMPANY_NAME")
	_ = os.Unsetenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	_ = os.Setenv("TEST", "1")
	main()
}

func TestMain_BadURL(t *testing.T) {
	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", "://bad-url")
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")
	_ = os.Setenv("TEST", "1")
	main()
}

func TestMain_ApplySuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/fieldValues" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"fieldValue":{"id":"fv1"}}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldTest := os.Getenv("TEST")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("TEST", oldTest)
	})
	os.Args = []string{"main", "-apply"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "t")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")
	_ = os.Setenv("TEST", "1")
	main()
}
