package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// captureRT records the incoming request and returns a configured response.
type captureRT struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (c *captureRT) RoundTrip(req *http.Request) (*http.Response, error) {
	c.req = req
	if c.resp != nil {
		return c.resp, c.err
	}
	// default empty 200
	return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{}`)), Header: make(http.Header)}, c.err
}

func makeResp(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}
}

func TestDo_GET_encodes_query_params(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{}}}

	rt := cc.HTTPClient.Transport.(*captureRT)

	_, err := cc.Do(context.Background(), "GET", "contacts", map[string]string{"email": "a@b.com"}, nil)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if rt.req == nil {
		t.Fatal("expected request recorded")
	}
	if rt.req.URL.RawQuery != "email=a%40b.com" {
		t.Fatalf("expected encoded query, got %q", rt.req.URL.RawQuery)
	}
}

func TestDo_POST_body_headers_and_debug(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	buf := &bytes.Buffer{}
	cc := &CoreClient{BaseURL: u, Token: "tok", HTTPClient: &http.Client{Transport: &captureRT{resp: makeResp(200, `{"ok":true}`)}}, Debug: true, DebugWriter: buf}

	rt := cc.HTTPClient.Transport.(*captureRT)

	payload := map[string]interface{}{"name": "x"}
	var out map[string]interface{}
	_, err := cc.Do(context.Background(), "POST", "customObjects/records/schema1", payload, &out)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if rt.req == nil {
		t.Fatal("expected request recorded")
	}
	// verify content-type
	if ct := rt.req.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected content-type application/json, got %q", ct)
	}
	// verify Api-Token header
	if at := rt.req.Header.Get("Api-Token"); at != "tok" {
		t.Fatalf("expected Api-Token header, got %q", at)
	}
	// verify request body was sent
	b, _ := io.ReadAll(rt.req.Body)
	var sent map[string]interface{}
	_ = json.Unmarshal(b, &sent)
	if sent["name"] != "x" {
		t.Fatalf("unexpected request body: %s", string(b))
	}
	// debug writer should contain header and JSON
	if !strings.Contains(buf.String(), "DEBUG OUTGOING POST") || !strings.Contains(buf.String(), "\"name\":\"x\"") {
		t.Fatalf("debug output missing, got: %s", buf.String())
	}
}

func TestDo_HTTPClient_error(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{err: io.ErrUnexpectedEOF}}}

	_, err := cc.Do(context.Background(), "GET", "contacts", nil, nil)
	if err == nil {
		t.Fatalf("expected error from HTTP client")
	}
}

func TestDo_Non2xx_returns_APIError(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	body := `{"Message":"bad","StatusCode":422}`
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{resp: makeResp(422, body)}}}

	apiResp, err := cc.Do(context.Background(), "GET", "contacts", nil, nil)
	if err == nil {
		t.Fatalf("expected API error for non-2xx")
	}
	if apiResp == nil {
		t.Fatalf("expected apiResp returned even on error")
	}
	if ae, ok := err.(*APIError); ok {
		if ae.StatusCode != 422 {
			t.Fatalf("expected StatusCode 422 in APIError, got %d", ae.StatusCode)
		}
	} else {
		t.Fatalf("expected *APIError, got %T", err)
	}
}

func TestDo_Success_unmarshal_out(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	body := `{"record":{"id":"r1"}}`
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{resp: makeResp(200, body)}}}

	var out map[string]map[string]string
	apiResp, err := cc.Do(context.Background(), "GET", "contacts", nil, &out)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if apiResp.StatusCode != 200 {
		t.Fatalf("unexpected status: %d", apiResp.StatusCode)
	}
	if out == nil || out["record"]["id"] != "r1" {
		t.Fatalf("unexpected out value: %#v", out)
	}
}

func TestDo_InvalidPath_returns_error(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{}}}

	_, err := cc.Do(context.Background(), "GET", "%", nil, nil)
	if err == nil {
		t.Fatalf("expected error for invalid path")
	}
}

func TestSetDebugFilter_prevents_debug_when_false(t *testing.T) {
	u, _ := url.Parse("https://example.com/api/3/")
	buf := &bytes.Buffer{}
	cc := &CoreClient{BaseURL: u, Token: "", HTTPClient: &http.Client{Transport: &captureRT{resp: makeResp(200, `{"ok":true}`)}}, Debug: true, DebugWriter: buf}
	cc.SetDebugFilter(func(method, path string) bool { return false })

	payload := map[string]interface{}{"a": "b"}
	_, err := cc.Do(context.Background(), "POST", "contacts", payload, nil)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no debug output when filter returns false, got: %s", buf.String())
	}
}
