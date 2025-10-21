package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteAccount deletes a single account by ID.
//
// What & Why:
//
//	Permanently deletes an account (organization) record.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#delete-account
//
// Parameters:
//
//	ctx: context
//	id: account ID
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *service) DeleteAccount(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#delete-account")
	}
	path := "accounts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
