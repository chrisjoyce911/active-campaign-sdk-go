//go:build ignore

// Legacy top-level entrypoint (ignored). This file is intentionally left
// build-ignored during migration to the new v3 packages.

package active_campaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// httpClient defines an interface for an http.Client implementation so that alternative
// If you'd prefer not to pass in your API token to this package, you can implement httpClient and
// handle adding the Api-Token on your own. See examples/custom_client.go to get started.
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

const (
	headerApiToken    = "Api-Token"
	headerContentType = "Content-Type"
)

// A Client manages communication with the Active Campaign API.
type Client struct {
	// HTTP client used to communicate with the API.
	client httpClient

	// Base URL for API requests.
	baseURL *url.URL

	// Token for API requests.
	token string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Active Campaign API.
	Contacts     *ContactsService
	Tags         *TagsService
	CustomFields *CustomFieldsService
	Lists        *ListsService
}

type service struct {
	client *Client
}

// ClientOpts are used to build a new client. If desired, a custom httpClient can be passed in.
type ClientOpts struct {
	HttpClient httpClient
	BaseUrl    string
	Token      string
}

// Meta is embedded in the Response struct.
type Meta struct {
	Total string `json:"total"`
}

// Links is embedded in the Response struct.
type Links struct {
	Options   string `json:"options"`
	Relations string `json:"relations"`
}

type ErrorResponse struct {
	Response interface{}
}

// InvalidError Response struct 'The request could not be processed, usually due to a missing or invalid parameter.'
type InvalidError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Code   string `json:"code"`
	Error  string `json:"error"`
	Source struct {
		Pointer string `json:"pointer"`
	} `json:"source"`
}

// NewClient returns a new Active Campaign API client. httpClient is provided to allow a
// custom client in specialized cases.
// If a nil httpClient is provided, a new http.DefaultClient will be used.
func NewClient(opts *ClientOpts) (*Client, error) {
	var httpClient httpClient
	if opts.HttpClient != nil {
		httpClient = opts.HttpClient
	} else {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := url.Parse(opts.BaseUrl)
	if err != nil {
		return nil, err
	}
	//go:build ignore

	package active_campaign

	// Legacy root-level source file: intentionally ignored during the v3 migration.
	// This file is retained for reference but excluded from normal builds.
