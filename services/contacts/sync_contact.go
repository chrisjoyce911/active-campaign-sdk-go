package contacts

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// SyncContact synchronizes contact data.
//
// What & Why:
//
//	Syncs or upserts contact data into ActiveCampaign. Useful for one-off
//	batch syncs or integrations.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#sync-contact-data
//
// Parameters:
//
//	ctx: context for cancellation/timeouts
//	req: request payload (currently interface{})
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) SyncContact(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "contacts/sync"
	apiResp, err := s.client.Do(ctx, http.MethodPost, path, req, &out)
	return out, apiResp, err
}
