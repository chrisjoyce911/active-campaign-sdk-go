package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/contacts"
)

func TestRun_DryRunWithEnvContact(t *testing.T) {
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

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	oldCID := os.Getenv("ACTIVE_CONTACT_ID")
	t.Cleanup(func() {
		_ = os.Setenv("ACTIVE_CONTACT_ID", oldCID)
	})
	_ = os.Setenv("ACTIVE_CONTACT_ID", "c1")

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "", "", "", "", "", false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRun_DryRunWithSearchFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/contacts" && r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contacts":[{"id":"c1","email":"test@example.com"}]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "", "", "", "", "", false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain_ApplyModeUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contacts":[{"id":"123","email":"test@example.com"}]}`)
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

func TestRun_ErrorSearch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/contacts" {
			http.Error(w, "error", 500)
		}
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "", "", "", "", "", false)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRun_ErrorCreate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/3/contacts" && r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"contacts":[]}`)
		} else if r.URL.Path == "/api/3/contacts" && r.Method == "POST" {
			http.Error(w, "error", 500)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "John", "Doe", "555", "ACME", "rto1", "", "", "", "", "", true)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRun_ErrorUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contacts":[{"id":"123","email":"test@example.com"}]}`)
			}
		case "/api/3/fieldValues":
			http.Error(w, "error", 500)
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "", "", "", "f1", "", true)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRun_ErrorList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contacts":[{"id":"123","email":"test@example.com"}]}`)
			}
		case "/api/3/contactLists":
			http.Error(w, "error", 500)
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "l1", "", "", "", "", true)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRun_ErrorTag(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/3/contacts":
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"contacts":[{"id":"123","email":"test@example.com"}]}`)
			}
		case "/api/3/contactTags":
			http.Error(w, "error", 500)
		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	svc := contacts.NewRealService(core)

	err = Run(context.Background(), svc, "test@example.com", "", "", "", "", "", "", "t1", "", "", "", true)
	if err == nil {
		t.Fatal("expected error")
	}
}
