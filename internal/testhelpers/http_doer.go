package testhelpers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// HTTPDoer constructs an http.Request similar to CoreClient and uses an
// internal fake http.Client (RoundTripper) to return a canned response.
// It records the final http.Request and body so tests can assert headers,
// resolved URL, and full request JSON.
type HTTPDoer struct {
	BaseURL string // optional base URL used to resolve relative paths
	Token   string // optional Api-Token header

	// Response to return
	RespStatus int
	RespBody   []byte
	RespErr    error

	// Recorded
	LastRequest     *http.Request
	LastRequestBody []byte
	// Optional transport for testing or custom clients. If nil, a fakeRoundTripper is used.
	Transport http.RoundTripper
}

// fakeRoundTripper returns the canned response and records the request.
type fakeRoundTripper struct {
	parent *HTTPDoer
}

func (f *fakeRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// record
	f.parent.LastRequest = req
	if req.Body != nil {
		b, _ := io.ReadAll(req.Body)
		f.parent.LastRequestBody = b
		// restore body for any downstream readers
		req.Body = io.NopCloser(bytes.NewReader(b))
	}

	if f.parent.RespErr != nil {
		return nil, f.parent.RespErr
	}

	// build response
	resp := &http.Response{
		StatusCode: f.parent.RespStatus,
		Body:       io.NopCloser(bytes.NewReader(f.parent.RespBody)),
		Header:     make(http.Header),
	}
	return resp, nil
}

// Do implements client.Doer. It builds an http.Request, sets headers,
// and uses a fake http.Client to execute the request (so RoundTripper sees it).
func (h *HTTPDoer) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*client.APIResponse, error) {
	// Resolve URL
	var fullURL string
	if h.BaseURL == "" {
		fullURL = path
	} else {
		// join base and path safely
		base, err := url.Parse(h.BaseURL)
		if err != nil {
			return nil, err
		}
		rel, err := url.Parse(path)
		if err != nil {
			return nil, err
		}
		fullURL = base.ResolveReference(rel).String()
	}

	var bodyReader io.Reader
	// GET with map[string]string -> query params
	if method == http.MethodGet {
		if params, ok := v.(map[string]string); ok && params != nil {
			u, _ := url.Parse(fullURL)
			q := u.Query()
			for k, val := range params {
				q.Set(k, val)
			}
			u.RawQuery = q.Encode()
			fullURL = u.String()
		}
	} else {
		if v != nil {
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if v != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if h.Token != "" {
		req.Header.Set("Api-Token", h.Token)
	}

	// set up client with provided transport or our fake roundtripper
	var tr http.RoundTripper
	if h.Transport != nil {
		tr = h.Transport
	} else {
		tr = &fakeRoundTripper{parent: h}
	}
	fake := &http.Client{Transport: tr}

	resp, err := fake.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResp := &client.APIResponse{HTTP: resp, Body: buf, StatusCode: resp.StatusCode}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr client.APIError
		_ = json.Unmarshal(buf, &apiErr)
		apiErr.StatusCode = resp.StatusCode
		apiErr.Body = buf
		return apiResp, &apiErr
	}
	if out != nil && len(buf) > 0 {
		if err := json.Unmarshal(buf, out); err != nil {
			return apiResp, err
		}
	}
	return apiResp, nil
}
