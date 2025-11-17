package main

import (
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
