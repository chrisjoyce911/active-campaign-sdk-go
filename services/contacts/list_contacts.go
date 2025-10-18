//go:build ignore

package contacts

import (
	"context"
	"fmt" // Ensure consistent fmt import

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListContacts lists contacts with optional filters.
//
// What & Why:
//
//	Retrieves a paginated list of contacts. Supports filters, sorting and
//	pagination parameters per the API.
//
// Docs:
//
//	Postman: https://www.postman.com/acdevrel/activecampaign-developer-relations/documentation/ju5a59q/activecampaign-api-v3
//	Reference: https://developers.activecampaign.com/reference#list-all-contacts
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	opts: map[string]string optional query params like limit, offset, email, etc.
//
// Returns:
//
//	(*ContactSearchResponse, *client.APIResponse, error)
//
// TODO: implement
func (s *service) ListContacts(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-all-contacts")
}
