package client

import (
	"net/http"
	"net/url"
)

// APIResponse wraps the low-level HTTP response and body for debugging and inspection.
type APIResponse struct {
	HTTP       *http.Response
	Body       []byte
	StatusCode int
}

// APIError represents a non-2xx response from the API.
type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
}

func (e *APIError) Error() string { return e.Message }

// BuildContactsSearchPath constructs the contacts search path for an email query.
func BuildContactsSearchPath(email string) string {
	return "/api/3/contacts?" + ("email=" + url.QueryEscape(email))
}
