package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactGoals returns goals for a contact.
//
// What & Why:
//
//	Retrieve goal completions for a contact for analytics or automation.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-goals
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactGoals(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/goals"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
