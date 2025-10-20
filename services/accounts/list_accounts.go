package accounts

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListAccounts lists accounts with optional filters.
//
// What & Why:
//
//	Returns a paginated list of accounts. Useful for admin views and syncs.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#list-accounts
//
// Parameters:
//
//	ctx: context
//	opts: query options
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) ListAccounts(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-accounts")
	}
	var out interface{}
	base := "accounts"
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
