package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactFieldValues returns custom field values for a contact.
//
// What & Why:
//
//	Retrieve custom field values associated with a contact. Useful for
//	displaying or syncing contact-specific data.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-field-values
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactFieldValues(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/fieldValues"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
