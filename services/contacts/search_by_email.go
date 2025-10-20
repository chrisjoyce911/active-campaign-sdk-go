package contacts

import (
	"context"
	"net/http"

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
//	(*ContactSearchResponse, *client.APIResponse, error)
//
// TODO: implement
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
//	(*ContactSearchResponse, *client.APIResponse, error)
func (s *RealService) SearchByEmail(ctx context.Context, email string) (*ContactSearchResponse, *client.APIResponse, error) {
	out := &ContactSearchResponse{}
	path := client.BuildContactsSearchPath(email)
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}
