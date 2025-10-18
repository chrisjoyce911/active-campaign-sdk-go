//go:build ignore

package tracking

import "context"

// TrackEvent sends an event to Site & Event Tracking.
//
// Docs: https://developers.activecampaign.com/reference#site-and-event-tracking
func (s *service) TrackEvent(ctx context.Context, payload interface{}) (*client.APIResponse, error) {
	return nil, nil
}
