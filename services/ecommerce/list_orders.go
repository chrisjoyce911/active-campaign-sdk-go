//go:build ignore

package ecommerce

import (
	"context"
	"fmt"
)

// ListOrders lists e-commerce orders.
func (s *service) ListOrders(ctx context.Context, opts map[string]string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#ecommerce")
}
