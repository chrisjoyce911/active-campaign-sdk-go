package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactOrganization returns the organization associated with a contact.
//
// What & Why:
//
//	Retrieves the organization for a contact useful for account-level
//	integrations or display.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-organization
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactOrganization(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/organization"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
