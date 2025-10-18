//go:build ignore

package deals

import (
	"context"
	"fmt"
)

// DeleteDeal deletes a deal.
func (s *service) DeleteDeal(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#deals")
}
