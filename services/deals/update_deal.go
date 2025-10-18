//go:build ignore

package deals

import (
	"context"
	"fmt"
)

// UpdateDeal updates a deal.
func (s *service) UpdateDeal(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#deals")
}
