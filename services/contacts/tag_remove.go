package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TagRemove removes a tag from a contact.
//
// What & Why:
//
//	Removes a tag association from a contact.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#delete-contact-tag
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	contactTagID: the ID of the contactTag association to remove
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *RealService) TagRemove(ctx context.Context, contactTagID string) (*client.APIResponse, error) {
	path := "contactTags/" + contactTagID
	return s.client.Do(ctx, http.MethodDelete, path, nil, nil)
}
