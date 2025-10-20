package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateAccount creates a new account (organization).
//
// What & Why:
//
//	Creates an organization or account object representing a company.
//	Useful when mapping contacts to organizations.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-an-account-new
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: payload for account creation
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) CreateAccount(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-an-account-new")
	}
	var out interface{}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "accounts", req, &out)
	return out, apiResp, err
}
