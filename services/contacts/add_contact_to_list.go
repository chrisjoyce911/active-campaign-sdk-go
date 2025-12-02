package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// AddContactToList adds a contact to a list.
//
// What & Why:
//
//	Adds a contact to an ActiveCampaign list (subscribe action).
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#add-contact-to-list
//
// Parameters:
//
//	ctx: context
//	req: request payload
//
// Returns:
//
//	(*AddContactToListResponse, *client.APIResponse, error)
func (s *RealService) AddContactToList(ctx context.Context, req *AddContactToListPayload) (*AddContactToListResponse, *client.APIResponse, error) {
	var out AddContactToListResponse
	body := struct {
		ContactList *AddContactToListPayload `json:"contactList"`
	}{ContactList: req}
	path := "contactLists"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, body, &out)
	return &out, apiResp, err
}
