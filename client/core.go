package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CoreClient struct {
	BaseURL    *url.URL
	Token      string
	HTTPClient *http.Client
	Cookie     string // Optional Cookie header value
	// Debug enables verbose outgoing request logging when true.
	Debug bool
	// DebugWriter, if non-nil, is used as the destination for debug output.
	// If nil, the standard library logger is used.
	DebugWriter io.Writer
	// DebugFilter, if non-nil, is called with (method, path) and debug output
	// will only be emitted when it returns true. If nil, the client will emit
	// debug output for all outgoing requests when Debug is enabled. Use
	// SetDebugFilter to provide a predicate to restrict which calls are logged.
	DebugFilter func(method, path string) bool
}

func NewCoreClient(baseURL, token string) (*CoreClient, error) {
	// Normalize baseURL to include /api/3/ if missing and ensure trailing slash.
	if baseURL != "" && !strings.Contains(baseURL, "/api/") {
		baseURL = strings.TrimRight(baseURL, "/") + "/api/3/"
	} else if baseURL != "" {
		baseURL = strings.TrimRight(baseURL, "/") + "/"
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &CoreClient{BaseURL: u, Token: token, HTTPClient: &http.Client{Timeout: 15 * time.Second}}, nil
}

// SetDebug enables or disables debug logging and optionally sets the writer
// where debug output will be written. If writer is nil, the default logger
// will be used (log.Printf).
func (c *CoreClient) SetDebug(enabled bool, writer io.Writer) {
	c.Debug = enabled
	c.DebugWriter = writer
}

// SetDebugFilter sets an optional predicate that controls when debug output
// will be emitted. If filter is nil the client will emit debug output for
// all outgoing calls when Debug is enabled. Provide a non-nil filter to
// restrict logging to a subset of requests (for example, only POSTs to
// custom objects endpoints).
func (c *CoreClient) SetDebugFilter(filter func(method, path string) bool) {
	c.DebugFilter = filter
}

func (c *CoreClient) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*APIResponse, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	reqURL := c.BaseURL.ResolveReference(rel)

	var body io.Reader
	// If this is a GET request and v is a map[string]string, treat it as
	// query parameters instead of a JSON body. Many callers pass opts as
	// map[string]string expecting query encoding.
	if strings.ToUpper(method) == http.MethodGet {
		if params, ok := v.(map[string]string); ok && params != nil {
			q := reqURL.Query()
			for k, val := range params {
				q.Set(k, val)
			}
			reqURL.RawQuery = q.Encode()
		}
	} else {
		if v != nil {
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if v != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Token != "" {
		req.Header.Set("Api-Token", c.Token)
	}
	if c.Cookie != "" {
		req.Header.Set("Cookie", c.Cookie)
	}

	// Conditional debug output controlled by Debug and optional DebugFilter
	shouldDebug := false
	if c.Debug {
		if c.DebugFilter != nil {
			// If a filter is provided, use it to decide when to emit debug output.
			shouldDebug = c.DebugFilter(method, reqURL.Path)
		} else {
			// No filter provided: show all calls by default when Debug is enabled.
			shouldDebug = true
		}
	}
	if shouldDebug {
		// prefer configured writer
		if c.DebugWriter != nil {
			_, _ = c.DebugWriter.Write([]byte("DEBUG OUTGOING " + method + " " + reqURL.String() + "\n"))
			for k, v := range req.Header {
				_, _ = c.DebugWriter.Write([]byte(k + ": " + strings.Join(v, ", ") + "\n"))
			}
			if v != nil {
				if bb, err := json.Marshal(v); err == nil {
					_, _ = c.DebugWriter.Write([]byte("body:\n"))
					_, _ = c.DebugWriter.Write(bb)
					_, _ = c.DebugWriter.Write([]byte("\n"))
				}
			}
			_, _ = c.DebugWriter.Write([]byte("\n"))
		} else {
			log.Printf("DEBUG OUTGOING %s %s", method, reqURL.String())
			for k, v := range req.Header {
				log.Printf("%s: %s", k, strings.Join(v, ", "))
			}
			if v != nil {
				if bb, err := json.Marshal(v); err == nil {
					log.Printf("body:\n%s", string(bb))
				}
			}
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResp := &APIResponse{HTTP: resp, Body: buf, StatusCode: resp.StatusCode, RetryAfter: resp.Header.Get("Retry-After")}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var apiErr APIError
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
