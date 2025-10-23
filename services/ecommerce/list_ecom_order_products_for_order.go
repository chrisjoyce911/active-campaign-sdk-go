package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListEcomOrderProductsForOrder lists order products for a specific order.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) ListEcomOrderProductsForOrder(ctx context.Context, orderID string, opts map[string]string) (*EcomOrderProductsListResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	var wrapper EcomOrderProductsListResponse
	path := "ecomOrders/" + orderID + "/orderProducts"
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, opts, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper, apiResp, nil
}
