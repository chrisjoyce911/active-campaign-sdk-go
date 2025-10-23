package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// ListCustomers lists e-commerce customers.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) ListCustomers(ctx context.Context, opts map[string]string) (*EcomCustomerListResponse, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	var wrapper EcomCustomerListResponse
	apiResp, err := s.client.Do(ctx, http.MethodGet, "ecomCustomers", opts, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper, apiResp, nil
}
