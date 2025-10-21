package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactAccountContacts returns the other contacts on the same account.
//
// What & Why:
//
//	List contacts that belong to the same account as the provided contact ID.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-account-contacts
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactAccountContacts(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/accountContacts"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
