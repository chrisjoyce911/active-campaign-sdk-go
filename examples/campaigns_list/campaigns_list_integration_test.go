//go:build integration
// +build integration

package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestExampleCampaignsList_Run builds the example binary and runs it against a
// local httptest.Server to verify it executes and prints campaign info. This
// is an integration test and is only run when `go test -tags=integration` is
// provided.
func TestExampleCampaignsList_Run(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/3/campaigns" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"campaigns":[{"id":"c1","name":"T","status":"1"}],"meta":{"total":"1"}}`)
	}))
	defer ts.Close()

	dir := "."
	outPath := filepath.Join(os.TempDir(), "campaigns_list_example_integration")
	cmd := exec.Command("go", "build", "-o", outPath, ".")
	cmd.Dir = dir
	if b, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v; output: %s", err, string(b))
	}
	defer os.Remove(outPath)

	run := exec.Command(outPath)
	run.Env = append(os.Environ(), "ACTIVE_URL="+ts.URL, "ACTIVE_TOKEN=test-token")
	outb, err := run.CombinedOutput()
	if err != nil {
		t.Fatalf("example run failed: %v; output: %s", err, string(outb))
	}
	if !strings.Contains(string(outb), "c1") {
		t.Fatalf("unexpected output: %s", string(outb))
	}
}
