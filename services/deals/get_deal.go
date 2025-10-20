package deals

import (
	"context"
	"net/http"

	"github.com/chrisjoyce911/active-campaign-sdk-go/client"
)

// GetDeal retrieves a deal by ID.
//
// What & Why:
//
//	Fetch a single deal resource by its identifier.
//
// Docs:
//
//	Reference: https://developers.activecampaign.com/reference#get-a-deal
//
// Parameters:
//
//	ctx: context
//	id: deal ID
//
// Returns:
//
//	(interface{}, *client.APIResponse, error)
func (s *RealService) GetDeal(ctx context.Context, id string) (interface{}, *client.APIResponse, error) {
	var out interface{}
	path := "deals/" + id
	apiResp, err := s.client.Do(ctx, http.MethodGet, path, nil, &out)
	return out, apiResp, err
}
