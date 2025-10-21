package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactLogs returns activity logs for a contact.
//
// What & Why:
//
//	Provide audit and activity logs for a contact to help debugging and
//	analytics.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-logs
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/logs"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
