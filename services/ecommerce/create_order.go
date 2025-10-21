package ecommerce

import (
	"context"
	"fmt"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateOrder creates an e-commerce order via the REST E-Commerce API.
func (s *service) CreateOrder(ctx context.Context, req interface{}) (interface{}, *client.APIResponse, error) {
	return nil, nil, fmt.Errorf("not implemented: see https://developers.activecampaign.com/reference#ecommerce")
}
