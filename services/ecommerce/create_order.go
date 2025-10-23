package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// CreateOrder creates an e-commerce order via the REST E-Commerce API.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) CreateOrder(ctx context.Context, req CreateOrderRequest) (*EcomOrder, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	// API returns wrapper {"ecomOrder": { ... }}
	var wrapper struct {
		EcomOrder EcomOrder `json:"ecomOrder"`
	}
	apiResp, err := s.client.Do(ctx, http.MethodPost, "ecomOrders", req, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.EcomOrder, apiResp, nil
}
