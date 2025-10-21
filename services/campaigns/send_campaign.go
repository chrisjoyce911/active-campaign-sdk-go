package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// SendCampaign triggers the sending or scheduling of a campaign.
//
// Docs: https://developers.activecampaign.com/reference#send-campaign
func (s *service) SendCampaign(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#send-campaign")
}
