//go:build ignore

package lists

import (
	"context"
	"fmt"
)

// AddContactToList adds a contact to a list or updates their subscription status.
//
// What & Why:
//
//	Corresponds to contactLists endpoint; use to subscribe/unsubscribe contacts
//	to lists programmatically.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#update-list-status-for-contact
//
// Parameters:
//
//	ctx: context
//	req: payload
//
// Returns:
//
//	(*UpdateContactListStatusResponse, *client.APIResponse, error)
func (s *service) AddContactToList(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-list-status-for-contact")
}
