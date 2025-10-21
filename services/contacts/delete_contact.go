package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteContact removes a contact by ID.
//
// What & Why:
//
//	Permanently removes a contact from the account. Note: ActiveCampaign may
//	soft-delete or anonymize contacts in some cases. Check API behavior.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#delete-contact
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: string contact ID
//
// Returns:
//
//	(*client.APIResponse, error)
func (s *RealService) DeleteContact(ctx context.Context, id string) (*client.APIResponse, error) {
	path := "contacts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
