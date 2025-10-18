//go:build ignore

package contacts

//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateContact updates an existing contact.
//
// What & Why:
//
//	Use to modify contact fields (email, names, phone) and custom field values.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#update-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: string contact ID
//	req: *CreateContactRequest or an UpdateRequest structure
//
// Returns:
//
//	(*CreateContactResponse, *client.APIResponse, error)
//
// TODO: implement
func (s *service) UpdateContact(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-contact")
}
