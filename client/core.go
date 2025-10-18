package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CoreClient struct {
	BaseURL    *url.URL
	Token      string
	HTTPClient *http.Client
}

func NewCoreClient(baseURL, token string) (*CoreClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &CoreClient{BaseURL: u, Token: token, HTTPClient: &http.Client{Timeout: 15 * time.Second}}, nil
}

func (c *CoreClient) Do(ctx context.Context, method, path string, v interface{}, out interface{}) (*APIResponse, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	reqURL := c.BaseURL.ResolveReference(rel)

	var body io.Reader
	if v != nil {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(b)
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

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResp := &APIResponse{HTTP: resp, Body: buf, StatusCode: resp.StatusCode}
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
