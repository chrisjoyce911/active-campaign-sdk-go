package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactDeals returns deal records linked to a contact.
//
// What & Why:
//
//	Retrieve the deals associated with a contact for CRM views and sync.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-deals
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactDeals(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/deals"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
