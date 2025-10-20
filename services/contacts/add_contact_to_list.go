package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddContactToList adds a contact to a list.
//
// What & Why:
//
//	Adds a contact to an ActiveCampaign list (subscribe action).
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#add-contact-to-list
//
// Parameters:
//
//	ctx: context
//	req: request payload
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) AddContactToList(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contactLists"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return out, apiResp, err
}
