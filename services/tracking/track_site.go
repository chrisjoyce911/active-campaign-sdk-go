package tracking

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// TrackSite registers a site visit or tracking pixel event.
func (s *service) TrackSite(ctx context.Context, payload interface{}) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#site-and-event-tracking")
}
