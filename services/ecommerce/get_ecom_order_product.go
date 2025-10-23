package ecommerce

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetEcomOrderProduct retrieves an ecom order product by ID.
// See: https://developers.activecampaign.com/reference#ecommerce
func (s *service) GetEcomOrderProduct(ctx context.Context, id string) (*EcomOrderProduct, *client.APIResponse, error) {
	if s == nil || s.client == nil {
		return nil, nil, fmt.Errorf("ecommerce service not configured: missing client")
	}
	var wrapper struct {
		EcomOrderProduct EcomOrderProduct `json:"ecomOrderProduct"`
	}
	path := "ecomOrderProducts/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &wrapper)
	if err != nil {
		return nil, apiResp, err
	}
	return &wrapper.EcomOrderProduct, apiResp, nil
}
