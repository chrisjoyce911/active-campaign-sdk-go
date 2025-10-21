package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateContact updates an existing contact via the API.
//
// What & Why:
//
//	Updates the contact resource with the given payload. Implemented on
//	RealService to perform the HTTP PUT against the API.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#update-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//	req: update payload
//
// Returns:
//
//	(*CreateContactResponse, *client.APIResponse, error)
func (s *RealService) UpdateContact(ctx context.Context, id string, req *CreateContactRequest) (*CreateContactResponse, *client.APIResponse, error) {
	out := &CreateContactResponse{}
	path := "contacts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, out)
	return out, apiResp, err
}
