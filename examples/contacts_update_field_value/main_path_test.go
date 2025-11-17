package main

import (
	"os"
	"testing"
)

func TestMain_DryRun(t *testing.T) {
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldCID := os.Getenv("ACTIVE_CONTACTID")
	oldCompany := os.Getenv("CONTACT_COMPANY_NAME")
	oldField := os.Getenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("ACTIVE_CONTACTID", oldCID)
		_ = os.Setenv("CONTACT_COMPANY_NAME", oldCompany)
		_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", oldField)
	})
	_ = os.Setenv("ACTIVE_URL", "http://example")
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("ACTIVE_CONTACTID", "c1")
	_ = os.Setenv("CONTACT_COMPANY_NAME", "ACME")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "f1")

	main()
}
