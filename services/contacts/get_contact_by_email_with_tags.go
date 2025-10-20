package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactByEmailWithTags fetches a contact by email including its tags.
//
// What & Why:
//
//	Return the contact identified by email along with tag information. Useful
//	for synchronization or lookups.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-by-email
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	email: contact email address
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactByEmailWithTags(ctx context.Context, email string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := client.BuildContactsSearchPath(email)
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
