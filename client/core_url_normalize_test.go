package client

import (
	"net/url"
	"testing"
)

func TestNewCoreClient_NormalizeBaseURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		relPath  string
		expected string
	}{
		{
			name:     "base without api segment",
			input:    "https://example.com",
			relPath:  "contacts/1/contactTags",
			expected: "https://example.com/api/3/contacts/1/contactTags",
		},
		{
			name:     "base with api but no trailing slash",
			input:    "https://example.com/api/3",
			relPath:  "contacts/1/contactTags",
			expected: "https://example.com/api/3/contacts/1/contactTags",
		},
		{
			name:     "base with api and trailing slash",
			input:    "https://example.com/api/3/",
			relPath:  "contacts/1/contactTags",
			expected: "https://example.com/api/3/contacts/1/contactTags",
		},
		{
			name:     "base with trailing slash only",
			input:    "https://example.com/",
			relPath:  "contacts/1/contactTags",
			expected: "https://example.com/api/3/contacts/1/contactTags",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCoreClient(tt.input, "token")
			if err != nil {
				t.Fatalf("NewCoreClient returned error: %v", err)
			}
			// Resolve relative path using the client logic
			rel, err := url.Parse(tt.relPath)
			if err != nil {
				t.Fatalf("failed to parse relPath: %v", err)
			}
			got := c.BaseURL.ResolveReference(rel).String()
			if got != tt.expected {
				t.Fatalf("expected resolved URL %q, got %q", tt.expected, got)
			}

			// No runtime HTTP client calls here; test only verifies URL normalization/resolution.
		})
	}
}
