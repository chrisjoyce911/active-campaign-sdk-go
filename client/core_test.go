package client

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// roundTripperFunc lets tests inject a custom RoundTripper implementation.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// Bad is a type that fails to marshal to JSON for testing error paths.
type Bad struct{}

func (b Bad) MarshalJSON() ([]byte, error) { return nil, errors.New("boom") }

// badReadCloser simulates an io.ReadCloser that errors on Read.
type badReadCloser struct{}

func (badReadCloser) Read(p []byte) (int, error) { return 0, io.ErrUnexpectedEOF }
func (badReadCloser) Close() error               { return nil }

func TestNewCoreClient_NormalizeAndErrors(t *testing.T) {
	// no api path -> appended
	c, err := NewCoreClient("https://example.com", "")
	assert.NoError(t, err)
	assert.NotNil(t, c.BaseURL)
	assert.Equal(t, "https://example.com/api/3/", c.BaseURL.String())

	// has api -> ensure trailing slash
	c2, err := NewCoreClient("https://example.com/api/3", "tok")
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/api/3/", c2.BaseURL.String())

	// invalid URL should produce an error
	_, err = NewCoreClient(":", "")
	assert.Error(t, err)
}

func TestCore_Do_GET_encodes_query_params_custom(t *testing.T) {
	var seen *http.Request
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		seen = r
		body := `{"ok":true}`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})

	base, _ := url.Parse("https://api.test/api/3/")
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}}

	var out map[string]interface{}
	apiResp, err := cc.Do(context.Background(), http.MethodGet, "contacts", map[string]string{"email": "a@b.com"}, &out)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	// ensure query encoded
	assert.NotNil(t, seen)
	assert.Equal(t, "GET", seen.Method)
	assert.Contains(t, seen.URL.RawQuery, "email=a%40b.com")
}

func TestCore_Do_POST_body_and_debug_custom(t *testing.T) {
	var seenBody []byte
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		b, _ := io.ReadAll(r.Body)
		seenBody = b
		body := `{"result":"ok"}`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})

	base, _ := url.Parse("https://api.test/api/3/")
	buf := &bytes.Buffer{}
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}, Token: "T", Debug: true, DebugWriter: buf}

	req := map[string]string{"name": "x"}
	var out map[string]string
	apiResp, err := cc.Do(context.Background(), http.MethodPost, "tags", req, &out)
	assert.NoError(t, err)
	assert.NotNil(t, apiResp)
	// body was JSON
	assert.Contains(t, string(seenBody), "name")
	// debug should have written something
	assert.Contains(t, buf.String(), "DEBUG OUTGOING")
}

func TestCore_Do_DebugFilter_blocks_output_custom(t *testing.T) {
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		body := `{"ok":true}`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})
	base, _ := url.Parse("https://api.test/api/3/")
	buf := &bytes.Buffer{}
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}, Debug: true, DebugWriter: buf}
	cc.SetDebugFilter(func(method, path string) bool { return false })

	var out map[string]interface{}
	_, err := cc.Do(context.Background(), http.MethodGet, "contacts", nil, &out)
	assert.NoError(t, err)
	// buffer should be empty due to filter
	assert.Equal(t, "", buf.String())
}

func TestCore_Do_HTTPClient_error_and_non2xx_custom(t *testing.T) {
	// simulate transport error
	rtErr := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return nil, io.ErrUnexpectedEOF
	})
	base, _ := url.Parse("https://api.test/api/3/")
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rtErr}}
	_, err := cc.Do(context.Background(), http.MethodGet, "contacts", nil, nil)
	assert.Error(t, err)

	// simulate 400 with JSON API error
	rtBad := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		body := `{"Message":"bad request"}`
		return &http.Response{StatusCode: 400, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})
	cc2 := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rtBad}}
	apiResp, err := cc2.Do(context.Background(), http.MethodGet, "contacts", nil, nil)
	assert.Error(t, err)
	// should be APIError
	if apiErr, ok := err.(*APIError); ok {
		assert.Equal(t, 400, apiErr.StatusCode)
		assert.Contains(t, string(apiErr.Body), "bad request")
	} else {
		t.Fatalf("expected APIError, got %T", err)
	}
	assert.NotNil(t, apiResp)
}

func TestCore_Do_unmarshal_errors_and_empty_body_custom(t *testing.T) {
	// invalid JSON body
	rtInvalid := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		body := `not-json`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})
	base, _ := url.Parse("https://api.test/api/3/")
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rtInvalid}}
	var out map[string]interface{}
	_, err := cc.Do(context.Background(), http.MethodGet, "contacts", nil, &out)
	assert.Error(t, err)

	// empty body should not error when out != nil
	rtEmpty := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(""))), Header: make(http.Header)}, nil
	})
	cc2 := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rtEmpty}}
	var out2 map[string]interface{}
	_, err = cc2.Do(context.Background(), http.MethodGet, "contacts", nil, &out2)
	assert.NoError(t, err)
}

func TestCore_Do_MarshalError(t *testing.T) {
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		// should not be reached when marshal fails
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(`{}`))), Header: make(http.Header)}, nil
	})
	base, _ := url.Parse("https://api.test/api/3/")
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}}

	// use package-level Bad which returns a marshal error
	_, err := cc.Do(context.Background(), http.MethodPost, "tags", Bad{}, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "boom")
}

func TestCore_Do_ResponseReadError(t *testing.T) {
	// response body that errors on read
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: badReadCloser{}, Header: make(http.Header)}, nil
	})
	base, _ := url.Parse("https://api.test/api/3/")
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}}

	_, err := cc.Do(context.Background(), http.MethodGet, "contacts", nil, nil)
	assert.Error(t, err)
}

func TestCore_Do_DebugUsesLogWhenWriterNil(t *testing.T) {
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		body := `{"ok":true}`
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(body))), Header: make(http.Header)}, nil
	})
	base, _ := url.Parse("https://api.test/api/3/")
	// capture standard logger output
	buf := &bytes.Buffer{}
	old := log.Writer()
	log.SetOutput(buf)
	defer log.SetOutput(old)

	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}, Debug: true}
	// provide v to generate debug body output
	_, err := cc.Do(context.Background(), http.MethodPost, "tags", map[string]string{"a": "b"}, nil)
	assert.NoError(t, err)
	// log output should contain our debug header
	assert.Contains(t, buf.String(), "DEBUG OUTGOING")
}

func TestCore_Do_NewRequestError(t *testing.T) {
	base, _ := url.Parse("https://api.test/api/3/")
	rt := roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte(`{}`))), Header: make(http.Header)}, nil
	})
	cc := &CoreClient{BaseURL: base, HTTPClient: &http.Client{Transport: rt}}

	// try methods with invalid characters which should be rejected by net/http
	invalidMethods := []string{"G ET", "GET\n", "GET\x00"}
	var gotErr bool
	for _, m := range invalidMethods {
		_, err := cc.Do(context.Background(), m, "contacts", nil, nil)
		if err != nil {
			gotErr = true
			break
		}
	}
	assert.True(t, gotErr, "expected at least one invalid method to cause an error")
}
