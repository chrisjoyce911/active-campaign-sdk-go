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
		if r.URL.Path != "/api/3/deals" {
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
		if !strings.Contains(q, "filters[group]=2") || !strings.Contains(q, "filters[stage]=7") {
			t.Fatalf("expected pipeline and stage filters, got: %s", q)
		}
		w.Header().Set("Content-Type", "application/json")
		// Since the example uses ListDealsAll, the helper will stop after the first
		// page when meta.total equals the number of returned items.
		_, _ = io.WriteString(w, `{"deals":[{"id":"46","title":"Able Hyena","group":"2","stage":"7"},{"id":"1","title":"Test Deal","group":"2","stage":"7"}],"meta":{"total":2}}`)
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
	if !strings.Contains(out, "deal 46") || !strings.Contains(out, "deal 1") {
		t.Fatalf("unexpected output: %s", out)
	}
	if !called {
		t.Fatalf("test server was not called")
	}
}
