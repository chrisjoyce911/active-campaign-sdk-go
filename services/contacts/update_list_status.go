//go:build ignore

package contacts

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateListStatusForContact updates a contact's subscription status for a list.
//
// What & Why:
//
//	Adds/removes a contact from a list or changes their subscription status.
//	This endpoint is useful when activating signup flows or unsubscribes.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#update-list-status-for-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: *UpdateListStatusForContactRequest
//
// Returns:
//
//	(*UpdateContactListStatusResponse, *client.APIResponse, error)
//
// TODO: implement
func (s *service) UpdateListStatusForContact(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#contacts")
}
func (s *service) UpdateListStatus(ctx context.Context, req *UpdateListStatusForContactRequest) (*UpdateContactListStatusResponse, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#contacts")
}
}
