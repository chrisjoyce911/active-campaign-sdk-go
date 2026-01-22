package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetContactBounceLogs returns bounce logs for a contact.
//
// What & Why:
//
//	Expose historic bounce events so callers can analyze deliverability.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-contact-bounce-logs
//
// Parameters:
//
//	ctx: context
//	contactID: internal contact id
//
// Returns:
//
//	(*BounceLogsResponse, *client.APIResponse, error)
func (s *RealService) GetContactBounceLogs(ctx context.Context, contactID string) (*BounceLogsResponse, *client.APIResponse, error) {
	var out BounceLogsResponse
	path := "contacts/" + contactID + "/bounceLogs"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return &out, apiResp, err
}
