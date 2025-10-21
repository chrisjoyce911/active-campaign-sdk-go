//go:build ignore

package ecommerce

import (
	"context"
	"fmt"
)

// GetOrder retrieves an e-commerce order by ID.
func (s *service) GetOrder(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#ecommerce")
}
