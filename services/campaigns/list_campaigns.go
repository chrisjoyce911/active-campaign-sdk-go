//go:build ignore

package campaigns

import (
	"context"
	"fmt"
)

// ListCampaigns lists campaigns with optional filters.
//
// Parameters:
//
//	ctx: context
//	opts: query options
//
// Returns: (interface{}, *client.APIResponse, error)
func (s *service) ListCampaigns(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#list-campaigns")
}
