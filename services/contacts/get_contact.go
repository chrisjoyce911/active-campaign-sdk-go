//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContact retrieves a contact by ID.
//
// What & Why:
//
//	Fetches the contact resource by its unique ID. Useful for inspecting full
//	contact details, links, and custom field values.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#get-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: string contact ID
//
// Returns:
//
//	(*CreatedContact, *acclient.APIResponse, error)
//
// TODO: implement
func (s *service) GetContact(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#get-contact")
}
