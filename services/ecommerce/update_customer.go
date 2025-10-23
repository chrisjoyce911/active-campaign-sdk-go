package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// UpdateCustomer updates an e-commerce customer.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) UpdateCustomer(ctx context.Context, id string, req UpdateCustomerRequest) (*EcomCustomer, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	var wrapper struct {
		EcomCustomer EcomCustomer `json:"ecomCustomer"`
	}
	path := "ecomCustomers/" + id
	apiResp, err := s.client.Do(ctx, http.MethodPut, path, req, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.EcomCustomer, apiResp, nil
}
