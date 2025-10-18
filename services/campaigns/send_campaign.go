//go:build ignore

package campaigns

import (
	"context"
	"fmt"
)

// SendCampaign triggers the sending or scheduling of a campaign.
//
// Docs: https://developers.activecampaign.com/reference#send-campaign
func (s *service) SendCampaign(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#send-campaign")
}
