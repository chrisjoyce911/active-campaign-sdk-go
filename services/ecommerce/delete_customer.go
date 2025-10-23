package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// DeleteCustomer deletes an e-commerce customer.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) DeleteCustomer(ctx context.Context, id string) (*client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	path := "ecomCustomers/" + id
	apiResp, err := s.client.Do(ctx, http.MethodDelete, path, nil, nil)
	return apiResp, err
}
