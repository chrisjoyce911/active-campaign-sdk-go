package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteCampaign deletes a campaign by ID.
func (s *service) DeleteCampaign(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#delete-campaign")
}
