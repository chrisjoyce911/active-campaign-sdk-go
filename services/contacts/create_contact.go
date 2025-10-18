//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateContact creates a new contact in ActiveCampaign.
//
// What & Why:
//
//	Creates a new contact resource. Use this when you want to add a contact to
//	the account. The endpoint supports nested custom field values and tags.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#create-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: *CreateContactRequest containing the contact payload
//
// Returns:
//
//	(*CreateContactResponse, *client.APIResponse, error)
//
// TODO:
//   - Implement HTTP call using client.CoreClient
//   - Unit tests and example
func (s *service) CreateContact(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#create-contact")
}
