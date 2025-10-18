//go:build ignore

package accounts

import (
	"context"
	"fmt"
)

// CreateAccount creates an account (organization) record in ActiveCampaign.
//
// What & Why:
//
//	Creates an organization or account object representing a company. Useful
//	when mapping contacts to organizations.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#create-account
//
// Parameters:
//
//	ctx: context
//	req: payload
//
// Returns:
//
//	(*CreateAccountResponse, *client.APIResponse, error)
func (s *service) CreateAccount(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference")
}
