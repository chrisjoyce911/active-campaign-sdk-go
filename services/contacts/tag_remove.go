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
//	Removes a tag association from a contact by first looking up the contact's tags
//	to find the association ID, then deleting the association.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#delete-contact-tag
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	contactID: the ID of the contact
//	tagID: the tag ID to remove
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *RealService) TagRemove(ctx context.Context, contactID, tagID string) (*client.APIResponse, error) {
	// First, get the contact's tags to find the association ID
	tagsResp, apiResp, err := s.TagsGet(ctx, contactID)
	if err != nil {
		return apiResp, err
	}
	if tagsResp == nil || tagsResp.ContactTags == nil {
		return apiResp, nil // No tags, nothing to remove
	}

	// Find the tag association
	var contactTagID string
	for _, ct := range *tagsResp.ContactTags {
		if ct.Tag == tagID {
			contactTagID = ct.ID
			break
		}
	}
	if contactTagID == "" {
		// Tag not found on contact
		return &client.APIResponse{StatusCode: 404}, nil // Or return an error?
	}

	// Now delete the association
	path := "contactTags/" + contactTagID
	return s.client.Do(ctx, http.MethodDelete, path, nil, nil)
}
