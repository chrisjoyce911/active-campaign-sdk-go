package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateAccount updates an account by ID.
//
// What & Why:
//
//	Update account properties such as name, address, or custom fields.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-account
//
// Parameters:
//
//	ctx: context
//	id: account ID
//	req: update payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *service) UpdateAccount(ctx context.Context, id string, req *UpdateAccountRequest) (*CreateAccountResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-account")
	}
	var out CreateAccountResponse
	path := "accounts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &out)
	if err != nil {
		return nil, apiResp, err
	}
	return &out, apiResp, nil
}
