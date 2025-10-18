//go:build ignore

package deals

import (
	"context"
	"fmt"
)

// ListDeals lists deals.
func (s *service) ListDeals(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#deals")
}
