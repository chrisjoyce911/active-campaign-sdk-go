package campaigns

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListCampaigns lists campaigns with optional filters.
//
// Parameters:
//
//	ctx: context
//	opts: query options
//
// Returns: (interface{}, *client.APIResponse, error)
func (s *service) ListCampaigns(ctx context.Context, opts interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-campaigns")
}
