package tracking

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TrackEvent sends an event to Site & Event Tracking.
//
// Docs: https://developers.activecampaign.com/reference#site-and-event-tracking
func (s *service) TrackEvent(ctx context.Context, payload interface{}) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#site-and-event-tracking")
}
