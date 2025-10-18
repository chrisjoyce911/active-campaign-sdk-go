//go:build ignore

package tracking

import (
	"context"
	"fmt"
)

// GetTrackingData retrieves tracking information for a contact or event.
func (s *service) GetTrackingData(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#site-and-event-tracking")
}
