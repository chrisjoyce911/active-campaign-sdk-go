package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListEcomOrderProducts lists ecom order products.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) ListEcomOrderProducts(ctx context.Context, opts map[string]string) (*EcomOrderProductsListResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	var wrapper EcomOrderProductsListResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "ecomOrderProducts", opts, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper, apiResp, nil
}
