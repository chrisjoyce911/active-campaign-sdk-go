package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

func TestMain_HappyPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/3/dealStages" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"dealStages":[{"id":"15","title":"Initial Contact","group":"2"}],"meta":{"total":1}}`)
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

func TestRun_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "server error", 500)
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "t")
	if err != nil {
		t.Fatal(err)
	}
	core.HTTPClient = ts.Client()
	svc := deals.NewRealService(core)
	var buf bytes.Buffer
	if err := Run(context.Background(), svc, &buf); err == nil {
		t.Fatalf("expected error")
	}
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
		t.Fatalf("expected non-zero exit")
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
	_ = os.Setenv("ACTIVE_TOKEN", "t")

	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
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
	_ = os.Setenv("ACTIVE_TOKEN", "t")

	main()
	if code == 0 {
		t.Fatalf("expected non-zero exit")
	}
}
