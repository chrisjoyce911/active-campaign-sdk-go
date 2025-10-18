//go:build ignore

package deals

import (
	"context"
	"fmt"
)

// CreateDeal creates a deal record.
func (s *service) CreateDeal(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#deals")
}
