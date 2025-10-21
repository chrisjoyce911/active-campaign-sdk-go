package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactDealList returns the list of deals associated with a contact.
//
// What & Why:
//
//	Expose a contact's deals so callers can present or sync relationship data.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-deal-list
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactDealList(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/deals"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
