package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactPlusAppend returns Plus Append data for a contact.
//
// What & Why:
//
//	Returns data produced by the Plus Append enrichment for the contact.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-plus-append
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactPlusAppend(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/plusAppend"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
