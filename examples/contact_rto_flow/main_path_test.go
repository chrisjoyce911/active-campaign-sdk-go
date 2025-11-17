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

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("CONTACT_EMAIL")
	oldCID := os.Getenv("ACTIVE_CONTACT_ID")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("CONTACT_EMAIL", oldEmail)
		_ = os.Setenv("ACTIVE_CONTACT_ID", oldCID)
	})
	os.Args = []string{"main"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("CONTACT_EMAIL", "test@example.com")
	_ = os.Setenv("ACTIVE_CONTACT_ID", "c1")

	main()
}

func TestMain_ApplyMode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contacts":[]}`)
			} else if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contact":{"id":"123"}}`)
			}
		case "/api/3/fieldValues":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"fieldValue":{"id":"456"}}`)
			}
		case "/api/3/contactLists":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contactList":{"id":"789"}}`)
			}
		case "/api/3/contactTags":
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contactTag":{"id":"101"}}`)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	oldArgs := os.Args
	oldURL := os.Getenv("ACTIVE_URL")
	oldTok := os.Getenv("ACTIVE_TOKEN")
	oldEmail := os.Getenv("CONTACT_EMAIL")
	oldFirst := os.Getenv("CONTACT_FIRST_NAME")
	oldLast := os.Getenv("CONTACT_LAST_NAME")
	oldList := os.Getenv("ACTIVE_CONTACT_LIST")
	oldTag := os.Getenv("ACTIVE_CONTACT_TAG")
	oldCFComp := os.Getenv("ACTIVE_CONTACT_CF_COMPANY_NAME")
	oldCFRTO := os.Getenv("ACTIVE_CONTACT_CF_RTO_ID")
	t.Cleanup(func() {
		os.Args = oldArgs
		_ = os.Setenv("ACTIVE_URL", oldURL)
		_ = os.Setenv("ACTIVE_TOKEN", oldTok)
		_ = os.Setenv("CONTACT_EMAIL", oldEmail)
		_ = os.Setenv("CONTACT_FIRST_NAME", oldFirst)
		_ = os.Setenv("CONTACT_LAST_NAME", oldLast)
		_ = os.Setenv("ACTIVE_CONTACT_LIST", oldList)
		_ = os.Setenv("ACTIVE_CONTACT_TAG", oldTag)
		_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", oldCFComp)
		_ = os.Setenv("ACTIVE_CONTACT_CF_RTO_ID", oldCFRTO)
	})
	os.Args = []string{"main", "-apply"}
	_ = os.Setenv("ACTIVE_URL", ts.URL)
	_ = os.Setenv("ACTIVE_TOKEN", "test-token")
	_ = os.Setenv("CONTACT_EMAIL", "test@example.com")
	_ = os.Setenv("CONTACT_FIRST_NAME", "John")
	_ = os.Setenv("CONTACT_LAST_NAME", "Doe")
	_ = os.Setenv("ACTIVE_CONTACT_LIST", "list1")
	_ = os.Setenv("ACTIVE_CONTACT_TAG", "tag1")
	_ = os.Setenv("ACTIVE_CONTACT_CF_COMPANY_NAME", "cf1")
	_ = os.Setenv("ACTIVE_CONTACT_CF_RTO_ID", "cf2")

	main()
}
