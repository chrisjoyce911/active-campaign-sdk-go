package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateCampaign updates an existing campaign.
func (s *service) UpdateCampaign(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#update-campaign")
}
