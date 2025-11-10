package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactLists returns lists that a contact belongs to.
//
// What & Why:
//
//	Retrieve the contact lists for a contact. Useful for determining
//	segmentation and subscription state.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-lists
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns a typed response wrapping the array under "contactLists".
func (s *RealService) GetContactLists(ctx context.Context, id string) (*ContactListsResponse, *client.APIResponse, error) {
	out := &ContactListsResponse{}
	path := "contacts/" + id + "/contactLists"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, out)
	return out, apiResp, err
}
