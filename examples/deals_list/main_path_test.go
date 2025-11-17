package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestMain_HappyPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/3/deals" {
			http.NotFound(w, r)
			return
		}
		// respond with two items and total=2 so pagination stops
		q, _ := url.QueryUnescape(r.URL.RawQuery)
		_ = q // ensure query is constructed; no assertion needed here
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"deals":[{"id":"1","title":"A","group":"2","stage":"7"},{"id":"2","title":"B","group":"2","stage":"7"}],"meta":{"total":2}}`)
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
