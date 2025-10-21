package deals

import (
	"context"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListDeals retrieves a page of deals.
//
// What & Why:
//
//	Return a list of deals with optional query parameters for filtering,
//	pagination, or sorting.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-all-deals
//
// Parameters:
//
//	ctx: context
//	opts: optional query parameters
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) ListDeals(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	base := "deals"
	if len(opts) > 0 {
		vals := url.Values{}
		for k, v := range opts {
			vals.Add(k, v)
		}
		base = base + "?" + vals.Encode()
	}
	apiResp, err := s.client.Do(ctx, http.MethodGet, base, nil, &out)
	return out, apiResp, err
}
