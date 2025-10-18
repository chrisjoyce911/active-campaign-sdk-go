//go:build ignore

package campaigns

import (
	"context"
	"fmt"
)

// DeleteCampaign deletes a campaign by ID.
func (s *service) DeleteCampaign(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#delete-campaign")
}
