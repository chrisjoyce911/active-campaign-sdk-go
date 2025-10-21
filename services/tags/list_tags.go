package tags

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListTags retrieves a page of tags with optional query params.
//
// It sends a GET request to /tags and accepts an optional map of
// query parameters (e.g., pagination or filtering keys). The response
// is unmarshalled into ListTagsResponse. When the service is not
// configured a not-implemented error is returned (useful during
// iterative migration and tests relying on zero-value receivers).
func (s *service) ListTags(ctx context.Context, opts map[string]string) (*ListTagsResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-tags")
	}
	var out ListTagsResponse
	base := "tags"
	if len(opts) > 0 {
		vals := url.Values{}
		for k, v := range opts {
			vals.Add(k, v)
		}
		base = base + "?" + vals.Encode()
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, base, nil, &out)
	return &out, apiResp, err
}
