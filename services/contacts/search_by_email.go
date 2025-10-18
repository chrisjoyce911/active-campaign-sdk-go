//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// SearchByEmail finds contacts matching the provided email.
//
// What & Why:
//
//	Use this to look up contacts by email. It handles URL encoding of special
//	characters (e.g., '+' signs) and returns matching contacts.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#list-all-contacts
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	email: email address to search
//
// Returns:
//
//	(*ContactSearchResponse, *acclient.APIResponse, error)
//
// TODO: implement
func (s *service) SearchByEmail(ctx context.Context, email string) (*ContactSearchResponse, *acclient.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#contacts")
}
