package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/chrisjoyce911/active-campaign-sdk-go/services/deals"
)

func TestRun_WithRealClientAndServer(t *testing.T) {
	called := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.URL.Path != "/api/3/dealStages" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.Header.Get("Api-Token") != "test-token" {
			t.Fatalf("missing or wrong Api-Token header: %q", r.Header.Get("Api-Token"))
		}
		// Verify encoded query contains our opts
		q, _ := url.QueryUnescape(r.URL.RawQuery)
		if !strings.Contains(q, "filters[d_groupid]=2") {
			t.Fatalf("expected pipeline filter in query, got: %s", q)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"dealStages":[{"id":"15","title":"Initial Contact","group":"4"},{"id":"16","title":"Qualifications - Low","group":"4"}],"meta":{"total":2}}`)
	}))
	defer ts.Close()

	core, err := client.NewCoreClient(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("NewCoreClient: %v", err)
	}
	core.HTTPClient = ts.Client()

	svc := deals.NewRealService(core)
	var buf bytes.Buffer
	if err := Run(context.Background(), svc, &buf); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "stage 15") || !strings.Contains(out, "stage 16") {
		t.Fatalf("unexpected output: %s", out)
	}
	if !called {
		t.Fatalf("test server was not called")
	}
}
