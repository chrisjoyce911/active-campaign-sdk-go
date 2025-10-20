package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactTrackingLogs returns tracking logs for a contact.
//
// What & Why:
//
//	Fetches tracking events (pageviews, events) associated with a contact.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-tracking-logs
//
// Parameters:
//
//	ctx: context
//	id: contact id
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetContactTrackingLogs(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/" + id + "/trackingLogs"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
