//go:build ignore

package ecommerce

import (
	"context"
	"fmt"
)

// UpdateOrder updates an e-commerce order.
func (s *service) UpdateOrder(ctx context.Context, id string, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#ecommerce")
}
