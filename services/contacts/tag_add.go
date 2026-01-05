package contacts

import (
	"context"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TagAdd adds a tag to an existing contact.
//
// What & Why:
//
//	Adds a tag to a contact by creating a contactTag association.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#create-contact-tag
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	contactID: the ID of the contact
//	tagID: the ID of the tag
//
// Returns:
//
//	(*ContactTagResponse, *client.APIResponse, error)
func (s *RealService) TagAdd(ctx context.Context, contactID, tagID string) (*ContactTagResponse, *client.APIResponse, error) {
	req := &ContactTagRequest{
		ContactTag: ContactTagPayload{
			Contact: contactID,
			Tag:     tagID,
		},
	}
	return s.CreateContactTag(ctx, req)
}