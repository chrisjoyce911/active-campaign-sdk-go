package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TagsGet returns tags for a contact.
//
// What & Why:
//
//	Fetches tags associated with a contact for display or sync operations.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-tags
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) TagsGet(ctx context.Context, id string) (*ContactTagsResponse, *client.APIResponse, error) {
	var out ContactTagsResponse
	// The ActiveCampaign API returns contact tags under /contacts/{id}/contactTags
	path := "contacts/" + id + "/contactTags"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
