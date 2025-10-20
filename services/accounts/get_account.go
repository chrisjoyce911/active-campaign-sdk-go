package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetAccount fetches a single account by ID.
//
// What & Why:
//
//	Fetches the organization (account) resource by its unique ID.
//	Useful for inspecting full account details and related contacts.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-account
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: account ID
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) GetAccount(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-account")
	}
	var out interface{}
	path := "accounts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
