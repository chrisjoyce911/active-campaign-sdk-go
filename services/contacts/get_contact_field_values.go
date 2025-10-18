//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactFieldValues fetches custom field values for a contact.
//
// What & Why:
//
//	Returns custom field values for a contact, useful when you need structured
//	custom data associated with a contact.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#get-contact-field-values
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	contactID: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
//
// # Removed stray import of acclient
//
// TODO: implement
func (s *service) GetContactFieldValues(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#contacts")
}
