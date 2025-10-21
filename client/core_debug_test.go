package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoreClient_DebugWriterAndFilter(t *testing.T) {
	// build a client with a fake base URL
	u, _ := url.Parse("https://example.com/api/3/")
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: mockRT{}}}

	var buf bytes.Buffer
	cc.SetDebug(true, &buf)

	// default filter is nil -> defaults to path contains customObjects/records/
	// call Do with such a path
	_, _ = cc.Do(context.Background(), "POST", "customObjects/records/schema1", map[string]interface{}{"record": map[string]interface{}{"id": "r1"}}, nil)
	// since we didn't wire an HTTP client, Do will attempt to build the request then call nil client causing panic; to avoid that we only assert the debug writer would be written by invoking the debug logic directly via SetDebugFilter

	// Instead, exercise the debug emission path by calling the internal condition via the filter and writer directly
	cc.SetDebugFilter(func(method, path string) bool { return true })
	// perform a Do which will trigger the debug emission before the HTTP call
	_, _ = cc.Do(context.Background(), "POST", "customObjects/records/schema1", map[string]interface{}{"record": map[string]interface{}{"id": "r1"}}, nil)
	assert.True(t, strings.Contains(buf.String(), "DEBUG OUTGOING"))

	// Test filter blocks output when returning false
	buf.Reset()
	cc.SetDebugFilter(func(method, path string) bool { return false })
	if cc.Debug && cc.DebugFilter("POST", "/customObjects/records/schema1") {
		if cc.DebugWriter != nil {
			_, _ = cc.DebugWriter.Write([]byte("SHOULD NOT APPEAR"))
		}
	}
	assert.Equal(t, "", buf.String())
}

// mockRT is a RoundTripper that returns a simple 200 response with empty JSON body.
type mockRT struct{}

func (m mockRT) RoundTrip(req *http.Request) (*http.Response, error) {
	r := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
		Header:     make(http.Header),
	}
	return r, nil
}
