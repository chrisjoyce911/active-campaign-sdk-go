package testhelpers

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestHTTPDoer_RoundTripReturnsError(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespErr: errors.New("net fail")}
	_, err := h.Do(context.Background(), "GET", "p", nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "net fail")
}

func TestHTTPDoer_BadBaseURL(t *testing.T) {
	h := &HTTPDoer{BaseURL: "://bad"}
	_, err := h.Do(context.Background(), "GET", "p", nil, nil)
	assert.Error(t, err)
}

func TestHTTPDoer_MarshalErrorForBody(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespStatus: 200, RespBody: []byte(`{"ok":true}`)}
	// channels cannot be JSON marshalled
	ch := make(chan int)
	_, err := h.Do(context.Background(), "POST", "p", ch, nil)
	assert.Error(t, err)
}

func TestHTTPDoer_Non2xxWithMalformedBodyReturnsAPIError(t *testing.T) {
	// non-JSON body should still return APIError but unmarshal will be ignored
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespStatus: 500, RespBody: []byte(`not-json`)}
	_, err := h.Do(context.Background(), "GET", "p", nil, nil)
	assert.Error(t, err)
	if apiErr, ok := err.(*client.APIError); ok {
		assert.Equal(t, 500, apiErr.StatusCode)
		assert.NotNil(t, apiErr.Body)
	}
}

func TestHTTPDoer_BadRelPath(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/"}
	// invalid relative path should cause url.Parse(path) to error
	_, err := h.Do(context.Background(), "GET", "\x00", nil, nil)
	assert.Error(t, err)
}

func TestHTTPDoer_SuccessUnmarshalOut(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespStatus: 200, RespBody: []byte(`{"x":1}`)}
	var out map[string]int
	resp, err := h.Do(context.Background(), "GET", "p", nil, &out)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if assert.NotNil(t, out) {
		// json numbers decode as float64 when into interface{}, but into map[string]int via Unmarshal into map[string]int should work
		// ensure key exists
		_, ok := out["x"]
		assert.True(t, ok)
	}
}

func TestHTTPDoer_NoBaseURL_BadPath_NewRequestError(t *testing.T) {
	h := &HTTPDoer{} // BaseURL empty -> fullURL == path
	_, err := h.Do(context.Background(), "GET", "\x00", nil, nil)
	assert.Error(t, err)
}

func TestHTTPDoer_UnmarshalOutFails(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", RespStatus: 200, RespBody: []byte(`{"x":"not-an-int"}`)}
	var out map[string]int
	resp, err := h.Do(context.Background(), "GET", "p", nil, &out)
	// Unmarshal into map[string]int should fail due to type mismatch
	assert.Error(t, err)
	// API response should still be returned
	assert.NotNil(t, resp)
}

type errReadCloser struct{}

func (e *errReadCloser) Read(p []byte) (int, error) { return 0, errors.New("read fail") }
func (e *errReadCloser) Close() error               { return nil }

type errRoundTripper struct{}

func (e *errRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: &errReadCloser{}}, nil
}

func TestHTTPDoer_BodyReadError(t *testing.T) {
	h := &HTTPDoer{BaseURL: "https://api.example.com/", Transport: &errRoundTripper{}}
	_, err := h.Do(context.Background(), "GET", "p", nil, nil)
	assert.Error(t, err)
}
