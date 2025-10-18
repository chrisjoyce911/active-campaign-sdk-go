//go:build ignore

package ecommerce

import (
	"context"
	"fmt"
)

// DeleteOrder deletes an order.
func (s *service) DeleteOrder(ctx context.Context, id string) (*client.APIResponse, error) {
	return nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#ecommerce")
}
