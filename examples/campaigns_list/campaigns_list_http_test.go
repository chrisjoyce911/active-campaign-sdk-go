package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/campaigns"
)

func TestRun_WithRealClientAndServer(t *testing.T) {
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.Path != "/api/3/campaigns" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("Api-Token") != "test-token" {
			t.Fatalf("missing or wrong Api-Token header: %q", r.Header.Get("Api-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"campaigns":[{"id":"c1","name":"One","status":"1"},{"id":"c2","name":"Two","status":"0"}],"meta":{"total":"2"}}`)
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("NewCoreClient: %v", err)
	}
	// Use the test server's client so the transport routes correctly.
	core.HTTPClient = ts.Client()

	svc := campaigns.NewRealService(core)
	var buf bytes.Buffer
	if err := Run(context.Background(), svc, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "c1") || !strings.Contains(out, "c2") {
		t.Fatalf("unexpected output: %s", out)
	}
	if !called {
		t.Fatalf("test server was not called")
	}
}
