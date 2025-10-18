//go:build ignore

package deals

import (
	"context"
	"fmt"
)

// GetDeal retrieves a deal by ID.
func (s *service) GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#deals")
}
